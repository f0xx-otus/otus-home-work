package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type field struct {
	structName, fieldName, fieldType, fieldTags string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Wrong argument amount")
	}
	filePath := os.Args[1]
	validateMap := make(map[string][]field)

	err := prepareValidateFields(filePath, &validateMap)
	if err != nil {
		log.Fatal("failed to prepare tags for validation", err)
	}

	for k, v := range validateMap {
		fmt.Println(" key and value", k, v)
	}

	path := filepath.Dir(filePath) + "/models_validators_generated.go"
	err = generate(path, validateMap)
	if err != nil {
		log.Fatal("failed to generate file", err)
	}
	fmt.Println(validateMap)
}

func prepareValidateFields(fPath string, valMap *map[string][]field) error {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, fPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, x := range file.Decls {
		if x, ok := x.(*ast.GenDecl); ok {
			if x.Tok != token.TYPE {
				continue
			}
			for _, x := range x.Specs {
				if x, ok := x.(*ast.TypeSpec); ok {
					structName := x.Name.Name
					if x, ok := x.Type.(*ast.StructType); ok {
						for _, x := range x.Fields.List {
							if len(x.Names) == 0 {
								continue
							}
							if x.Tag != nil {
								fieldName := x.Names[0].Name
								var buf bytes.Buffer
								err = format.Node(&buf, fileset, x.Type)
								if err != nil {
									log.Fatal("error")
									return err
								}
								fieldType := buf.String()
								tags := x.Tag.Value
								if strings.Contains(tags, "validate:") {
									tagSlice := strings.Split(tags, " ")
									for _, t := range tagSlice {
										if strings.Contains(t, "validate:") {
											tag := strings.Trim(t, "`")
											tag = strings.Trim(tag, "validate:")
											tag = strings.Trim(tag, "\"")
											tag = strings.ReplaceAll(tag, ":", " ")
											tag = strings.ReplaceAll(tag, ",", " ")
											tag = strings.ReplaceAll(tag, "|", " ")
											(*valMap)[structName] = append((*valMap)[structName], field{structName, fieldName, fieldType, tag})
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}
