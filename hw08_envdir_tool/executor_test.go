package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("non-zero exit code", func(t *testing.T) {
		inputPath := "testdata/testDir/"
		err := os.Mkdir(inputPath, 0777)
		file, err := os.Create(inputPath + "exitFile.sh")
		if err != nil {
			log.Fatal(err)
		}
		err = file.Chmod(0777)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString("#!/usr/bin/env bash\n exit 1")
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err = os.RemoveAll(inputPath)
			if err != nil {
				log.Fatal(err)
			}
		}()

		cmd := []string{inputPath + "exitFile.sh"}
		env := Environment{"FOO": "BAR"}
		exitCode := RunCmd(cmd, env)
		fmt.Println(exitCode)
		require.Equal(t, 1, exitCode)
	})
	t.Run("check arguments", func(t *testing.T) {
		inputPath := "testdata/testDir/"
		err := os.Mkdir(inputPath, 0777)
		file, err := os.Create(inputPath + "exitFile.sh")
		if err != nil {
			log.Fatal(err)
		}
		err = file.Chmod(0777)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString("#!/usr/bin/env bash\n let \"a = $1 + $2\"\n if [[ a -eq 3 ]]\n " +
			"then\n     exit 0\n else\n     exit 1\n fi")
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err = os.RemoveAll(inputPath)
			if err != nil {
				log.Fatal(err)
			}
		}()

		cmd := []string{inputPath + "exitFile.sh", "1", "2"}
		env := Environment{"FOO": "BAR"}
		exitCode := RunCmd(cmd, env)
		fmt.Println(exitCode)
		require.Equal(t, 0, exitCode)
	})
}
