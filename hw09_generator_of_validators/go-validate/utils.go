package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"os"
)

type userDefinedTypes struct {
	RawTypeName string
	IsArrayType bool
}

// collectNonStructTypeInfo find all non struct user defined types and collect them in userDefinedTypes struct map.
func collectNonStructTypeInfo(node *ast.File) map[string]userDefinedTypes {
	res := make(map[string]userDefinedTypes)
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

			_, ok = currType.Type.(*ast.StructType)
			if !ok {
				// Collect info about user defined types that are not a structs
				userDefType := userDefinedTypes{}
				switch currType := currType.Type.(type) {
				case *ast.Ident:
					userDefType.RawTypeName = currType.Name
				case *ast.ArrayType:
					userDefType.RawTypeName = currType.Elt.(*ast.Ident).Name
					userDefType.IsArrayType = true
				}

				res[currType.Name.Name] = userDefType
				continue
			}
		}
	}
	return res
}

// formatSourceAndWrite format final source code for validation functions and write it to result file.
func formatSourceAndWrite(resFileName string, buf *bytes.Buffer) error {
	out, err := os.Create(resFileName)
	if err != nil {
		return fmt.Errorf("error while create result file %w", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		os.Remove(resFileName)
		return err
	}
	_, err = out.Write(p)
	if err != nil {
		os.Remove(resFileName)
		return err
	}
	return nil
}
