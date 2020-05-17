package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

const (
	basicTemplate = `//DO NOT EDIT THIS FILE.
//Generated by go-validation

package {{ . }}

import (
	"errors"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

`
	validateFuncTemplate = `func (v {{ . }}) Validate() ([]ValidationError, error) {
var resError []ValidationError
var err error
`
	lenValidationTemlate = `	if len({{ .FieldName }}) != {{ .Val }} {
	resError = append(resError, ValidationError{Field: {{ .FieldName }}, Err: errors.New("field lenght validation error"),})
	}`

	validationByReflect = `		switch reflect.ValueOf({{ . }}).Kind() `
)

type LenValidation struct {
	FieldName string
	Val       string
}

var ErrUnsupportedFieldType error = errors.New("unsupported filed type")
var ErrUnsupportedTag error = errors.New("unsupported tag fro field type")

func main() {
	t := template.New("Generate Template")
	t.Parse(basicTemplate)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	resFileName := fmt.Sprintf("%s_validation.go", strings.TrimSuffix(os.Args[1], ".go"))
	out, err := os.Create(resFileName)
	if err != nil {
		panic(fmt.Errorf("error while create result file %w", err))
	}
	t.Execute(out, node.Name.Name)

	t.Parse(validateFuncTemplate)
	// generate validation
	for _, f := range node.Decls {
		g, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range g.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			currStruct, ok := currType.Type.(*ast.StructType)
			if !ok {
				continue
			}
			t.Execute(out, currType.Name.Name)

			for _, filed := range currStruct.Fields.List {
				if filed.Tag == nil {
					continue
				}
				tagString, ok := reflect.StructTag(strings.Trim(filed.Tag.Value, "`")).Lookup("validate")
				if !ok {
					continue
				}
				tagList := strings.Split(tagString, "|")
				fieldName := filed.Names[0].Name

				var fieldTypeName string
				var isSliceType bool
				switch filed.Type.(type) {
				case *ast.Ident:
					fieldTypeName = filed.Type.(*ast.Ident).Name
				case *ast.ArrayType:
					fieldTypeName = filed.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
					isSliceType = true
				default:
					panic(ErrUnsupportedFieldType)
				}
				resFieldName := "v." + fieldName
				if isSliceType {
					fmt.Fprintln(out, `	for _, val := range v.`+fieldName+` {`)
					resFieldName = "val"
				}
				err := processTagList(out, tagList, resFieldName, fieldTypeName)
				if err != nil {
					panic(fmt.Errorf("error while processing tags: %w", err))
				}
				if isSliceType {
					fmt.Fprintln(out, `	}`)
				}
			}

			fmt.Fprintln(out, `}`)
			fmt.Fprintln(out)
		}
	}
}

func processTagList(out *os.File, tagList []string, fieldName string, fieldTypeName string) error {
	t := template.New("Validation template")
	for _, tag := range tagList {
		switch fieldTypeName {
		case "string":
			switch {
			case strings.HasPrefix(tag, "len"):
				t.Parse(lenValidationTemlate)
				val := strings.TrimPrefix(tag, "len:")
				t.Execute(out, LenValidation{
					FieldName: fieldName,
					Val:       val,
				})
			case strings.HasPrefix(tag, "regexp"):
				log.Println("regexp validation")
			case strings.HasPrefix(tag, "in"):
				log.Println("in validation")
			default:
				return ErrUnsupportedTag
			}
		case "int":
		default:
			log.Println(fieldTypeName)
			return ErrUnsupportedFieldType
		}

		// switch {
		// case strings.HasPrefix(tag, "len"):
		// 	log.Println("len validation")
		// case strings.HasPrefix(tag, "min"):
		// 	log.Println("min validation")
		// case strings.HasPrefix(tag, "max"):
		// 	log.Println("max validation")
		// case strings.HasPrefix(tag, "regexp"):
		// 	log.Println("regexp validation")
		// case strings.HasPrefix(tag, "in"):
		// 	log.Println("in validation")
		// }
	}
	return nil
}
