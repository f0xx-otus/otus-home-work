package main

import (
	"log"
	"os"
	"path/filepath"
)

type item struct {
	structName, fieldName, fieldType, fieldTags string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Wrong argument amount")
	}
	filePath := os.Args[1]
	validateMap := make(map[string][]item)

	err := prepareValidateFields(filePath, &validateMap)
	if err != nil {
		log.Fatal("failed to prepare tags for validation", err)
	}

	path := filepath.Dir(filePath) + "/models_validators_generated.go"
	err = generate(path, validateMap)
	if err != nil {
		log.Fatal("failed to generate file", err)
	}
}
