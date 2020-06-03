package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

const (
	validateTagName     = "validate"
	tagSplitSymbol      = ":"
	inArraySplitSymbol  = ","
	lenValidationTag    = "len"
	inValidationTag     = "in"
	minValidationTag    = "min"
	maxValidationTag    = "max"
	regexpValidationTag = "regexp"
	basicFileNameSuffix = "validation.go"
	sourceFileExtension = ".go"
)

type validationParams struct {
	FieldName       string
	RawFieldName    string
	Val             string
	IsNeedToConvert bool
	RawFiledType    string
}

// ErrFormat for erorrs during final formating and writing source file.
var ErrFormat = "error while formatting final file"

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Wrong number of argumets. Two arguments require")
	}

	err := processFile(os.Args[1])
	if err != nil {
		log.Fatalf("erorr while generate validation func: %s", err)
	}
}

// processFile function process input file and generate result source.
func processFile(fileName string) error {
	fset := token.NewFileSet()
	var err error
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	t, err := prepareTemplate()
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(&buf, basicTemplateName, node.Name.Name)
	if err != nil {
		return fmt.Errorf("%s:%w", ErrExecuteTemplate, err)
	}

	// first we need to collect info about non struct user defined types
	externalUserNonStructTypes := collectNonStructTypeInfo(node)

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
			err := processStruct(&buf, t, currStruct, currType, externalUserNonStructTypes)
			if err != nil {
				return err
			}
		}
	}

	resFileName := fmt.Sprintf("%s_%s", strings.TrimSuffix(fileName, sourceFileExtension), basicFileNameSuffix)

	err = formatSourceAndWrite(resFileName, &buf)
	if err != nil {
		return fmt.Errorf("%s:%w", ErrFormat, err)
	}
	return nil
}
