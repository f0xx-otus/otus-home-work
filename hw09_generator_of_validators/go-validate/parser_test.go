package main

import (
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func TestEmptyTag(t *testing.T) {
	path := "/tmp/testData.go"
	testMap := make(map[string][]item)
	testData := `package models
type (
	User struct {
		Phones []string` + "`" + `validate:":11"` + "`" + `
	}
)`

	f, err := os.Create(path)
	check(err)

	_, err = f.WriteString(testData)
	check(err)

	err = f.Close()
	check(err)

	err = prepareValidateFields(path, &testMap)
	if err != nil {
		log.Fatal("failed to prepare tags for validation", err)
	}

	expectData := item{"User", "Phones", "[]string", " 11"}
	actualData := testMap["User"]

	require.Equal(t, expectData, actualData[0])
}

func TestWrongType(t *testing.T) {
	path := "/tmp/testData.go"
	testMap := make(map[string][]item)
	testData := `package models
type (
	User struct {
		Phones int64` + "`" + `validate:"asdfg:11"` + "`" + `
	}
)`

	f, err := os.Create(path)
	check(err)

	_, err = f.WriteString(testData)
	check(err)

	err = f.Close()
	check(err)

	err = prepareValidateFields(path, &testMap)

	require.NotNil(t, err)
}

func TestWrongFilePath(t *testing.T) {
	path := "/tmp/testData.go"
	testMap := make(map[string][]item)

	err := prepareValidateFields(path, &testMap)

	require.NotNil(t, err)
}
