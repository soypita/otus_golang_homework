package main

import (
	"fmt"
	"go/ast"
	"io"
	"reflect"
	"strings"
	"text/template"
)

// ErrUnsupportedFieldType for unsupported struct filed type.
var ErrUnsupportedFieldType = "unsupported filed type"

// ErrUnsupportedTag for unsupported tag.
var ErrUnsupportedTag = "unsupported tag for field type"

// processStruct process user defined structs and generate validation function if require.
func processStruct(buf io.Writer, t *template.Template, currStruct *ast.StructType, currType *ast.TypeSpec, externalUserNonStructTypes map[string]userDefinedTypes) error {
	var isFunGenerated bool
	for _, filed := range currStruct.Fields.List {
		if filed.Tag == nil {
			continue
		}
		tagString, ok := reflect.StructTag(strings.Trim(filed.Tag.Value, "`")).Lookup(validateTagName)
		if !ok {
			continue
		}

		// Need to check to handle multiple execution of func template
		if !isFunGenerated {
			err := t.ExecuteTemplate(buf, validateFuncTemplateName, currType.Name.Name)
			if err != nil {
				return err
			}
			isFunGenerated = true
		}

		err := handleTagString(buf, t, filed, tagString, externalUserNonStructTypes)
		if err != nil {
			return err
		}
	}
	if isFunGenerated {
		fmt.Fprintln(buf, `return resError, err }`)
	}
	return nil
}

// handleTagString process full tag string on struct field.
func handleTagString(buf io.Writer, t *template.Template, filed *ast.Field, tagString string, externalUserNonStructTypes map[string]userDefinedTypes) error {
	var fieldTypeName string
	tagList := strings.Split(tagString, "|")
	fieldName := filed.Names[0].Name

	var isSliceType, isNeedToConvert bool
	switch fieldType := filed.Type.(type) {
	case *ast.Ident:
		// lets check type name in external types map
		userType, ok := externalUserNonStructTypes[fieldType.Name]
		if !ok {
			fieldTypeName = fieldType.Name
		} else {
			fieldTypeName = userType.RawTypeName
			isNeedToConvert = true
		}

		isSliceType = userType.IsArrayType
	case *ast.ArrayType:
		// validate multi-dim array
		if _, ok := fieldType.Elt.(*ast.ArrayType); ok {
			return fmt.Errorf(ErrUnsupportedFieldType)
		}
		// lets check type name in external types map
		userType, ok := externalUserNonStructTypes[fieldType.Elt.(*ast.Ident).Name]
		if !ok {
			fieldTypeName = fieldType.Elt.(*ast.Ident).Name
		} else {
			fieldTypeName = userType.RawTypeName
			isNeedToConvert = true
		}

		// go-validate not support nested arrays
		if userType.IsArrayType {
			return fmt.Errorf(ErrUnsupportedFieldType)
		}
		isSliceType = true
	default:
		return fmt.Errorf(ErrUnsupportedFieldType)
	}

	resFieldName := "v." + fieldName
	if isSliceType {
		fmt.Fprintln(buf, `	for _, val := range v.`+fieldName+` {`)
		resFieldName = "val"
	}
	err := processTagList(t, buf, tagList, resFieldName, fieldTypeName, fieldName, isNeedToConvert)
	if err != nil {
		return err
	}
	if isSliceType {
		fmt.Fprintln(buf, `	}`)
	}
	return nil
}

// processTagList generate validation code for slice of tags on a struct field.
func processTagList(t *template.Template, out io.Writer, tagList []string, fieldName string, fieldTypeName string, rawFieldName string, isNeedToConvert bool) error {
	for _, tag := range tagList {
		switch fieldTypeName {
		case "string":
			err := processStrTags(out, fieldName, fieldTypeName, rawFieldName, isNeedToConvert, tag, t)
			if err != nil {
				return err
			}
		case "int":
			err := processIntTags(out, fieldName, fieldTypeName, rawFieldName, isNeedToConvert, tag, t)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf(ErrUnsupportedTag)
		}
	}
	return nil
}

// processStrTags generate validation code for string field type.
func processStrTags(out io.Writer, fieldName string, fieldTypeName string, rawFieldName string, isNeedToConvert bool, tag string, t *template.Template) error {
	switch {
	case strings.HasPrefix(tag, lenValidationTag):
		val := strings.TrimPrefix(tag, fmt.Sprintf("%s%s", lenValidationTag, tagSplitSymbol))
		err := t.ExecuteTemplate(out, lenValidationTemlateName, validationParams{
			FieldName:       fieldName,
			Val:             val,
			RawFieldName:    rawFieldName,
			IsNeedToConvert: isNeedToConvert,
			RawFiledType:    fieldTypeName,
		})
		if err != nil {
			return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
		}
	case strings.HasPrefix(tag, regexpValidationTag):
		val := strings.TrimPrefix(tag, fmt.Sprintf("%s%s", regexpValidationTag, tagSplitSymbol))
		err := t.ExecuteTemplate(out, regexpValidationTemplateName, validationParams{
			FieldName:       fieldName,
			Val:             fmt.Sprintf("`%s`", val),
			RawFieldName:    rawFieldName,
			IsNeedToConvert: isNeedToConvert,
			RawFiledType:    fieldTypeName,
		})
		if err != nil {
			return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
		}
	case strings.HasPrefix(tag, inValidationTag):
		val := strings.TrimPrefix(tag, fmt.Sprintf("%s%s", inValidationTag, tagSplitSymbol))
		splitVal := strings.Split(val, inArraySplitSymbol)
		for i, v := range splitVal {
			splitVal[i] = fmt.Sprintf(`"%s"`, v)
		}
		err := t.ExecuteTemplate(out, inListValidationTemplateName, validationParams{
			FieldName:       fieldName,
			Val:             strings.Join(splitVal, inArraySplitSymbol),
			RawFieldName:    rawFieldName,
			IsNeedToConvert: isNeedToConvert,
			RawFiledType:    fieldTypeName,
		})
		if err != nil {
			return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
		}
	default:
		return fmt.Errorf(ErrUnsupportedTag)
	}
	return nil
}

// processIntTags generate validation code for int field type.
func processIntTags(out io.Writer, fieldName string, fieldTypeName string, rawFieldName string, isNeedToConvert bool, tag string, t *template.Template) error {
	switch {
	case strings.HasPrefix(tag, minValidationTag):
		val := strings.TrimPrefix(tag, fmt.Sprintf("%s%s", minValidationTag, tagSplitSymbol))
		err := t.ExecuteTemplate(out, minValidationTemplateName, validationParams{
			FieldName:       fieldName,
			Val:             val,
			RawFieldName:    rawFieldName,
			IsNeedToConvert: isNeedToConvert,
			RawFiledType:    fieldTypeName,
		})
		if err != nil {
			return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
		}
	case strings.HasPrefix(tag, maxValidationTag):
		val := strings.TrimPrefix(tag, fmt.Sprintf("%s%s", maxValidationTag, tagSplitSymbol))
		err := t.ExecuteTemplate(out, maxValidationTemplateName, validationParams{
			FieldName:       fieldName,
			Val:             val,
			RawFieldName:    rawFieldName,
			IsNeedToConvert: isNeedToConvert,
			RawFiledType:    fieldTypeName,
		})
		if err != nil {
			return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
		}
	case strings.HasPrefix(tag, inValidationTag):
		val := strings.TrimPrefix(tag, fmt.Sprintf("%s%s", inValidationTag, tagSplitSymbol))
		err := t.ExecuteTemplate(out, inListValidationTemplateName, validationParams{
			FieldName:       fieldName,
			Val:             val,
			RawFieldName:    rawFieldName,
			IsNeedToConvert: isNeedToConvert,
			RawFiledType:    fieldTypeName,
		})
		if err != nil {
			return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
		}
	default:
		return fmt.Errorf(ErrUnsupportedTag)
	}
	return nil
}
