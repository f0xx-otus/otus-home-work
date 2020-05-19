package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	envList := make(Environment, len(files))
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fileName := f.Name()
		line := ""
		if f.Size() == 0 {
			envList[fileName] = ""
			break
		}
		file, err := os.Open(dir + "/" + fileName)
		if err != nil {
			return envList, err
		}
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		line = scanner.Text()
		line = strings.TrimRight(line, " ")
		line = strings.Replace(line, "\x00", "\n", -1)
		fileName = string(bytes.Replace([]byte(fileName), []byte("="), []byte(""), -1))
		envList[fileName] = line
		err = file.Close()
		if err != nil {
			log.Panic("can't close output file", err)
		}
	}
	return envList, nil
}
