package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

func prepareValidateFields(fPath string, valMap *map[string][]item) error {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, fPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			break
		}
		if genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				break
			}
			structName := typeSpec.Name.Name
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				break
			}
			for _, field := range structType.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				if field.Tag != nil {
					fieldName := field.Names[0].Name
					fieldType, err := convertType(fileset, fPath, field.Type)
					if err != nil {
						return err
					}

					tags := field.Tag.Value
					tag := prepareTags(tags)
					(*valMap)[structName] = append((*valMap)[structName], item{structName,
						fieldName, fieldType, tag})
				}
			}
		}
	}
	return nil
}

func convertType(fset *token.FileSet, path string, rawFieldType ast.Expr) (string, error) {
	normalTypes := []string{"int", "[]int", "string", "[]string"}
	var typeName bytes.Buffer
	_ = printer.Fprint(&typeName, fset, rawFieldType)
	fieldType := typeName.String()

	for _, v := range normalTypes {
		if fieldType == v {
			return v, nil
		}
	}
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("wrong argument: %w", err)
	}

	for _, f := range node.Decls {
		genD, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genD.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			_, ok = currType.Type.(*ast.StructType)
			if !ok {
				var typeName bytes.Buffer
				_ = printer.Fprint(&typeName, fset, currType.Type)
				return typeName.String(), nil
			}
		}
	}
	return "", errors.New("failed to convert type")
}

func prepareTags(tags string) string {
	tagSlice := strings.Split(tags, " ")
	for _, t := range tagSlice {
		if strings.Contains(t, "validate") {
			tag := strings.Trim(t, "`")
			tag = strings.Trim(tag, "validte:")
			tag = strings.Trim(tag, "\"")
			tag = strings.ReplaceAll(tag, ":", " ")
			tag = strings.ReplaceAll(tag, ",", " ")
			tag = strings.ReplaceAll(tag, "|", " ")
			return tag
		}
	}
	return ""
}
