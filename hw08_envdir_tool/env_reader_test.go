package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("directory without files", func(t *testing.T) {
		var envSlice []string
		inputPath := "testdata/emptyDir"
		err := os.Mkdir(inputPath, 0777)
		defer func() {
			err = os.Remove(inputPath)
			if err != nil {
				log.Fatal(err)
			}
		}()
		if err != nil {
			log.Fatal(err)
		}
		env, err := ReadDir(inputPath)
		if err != nil {
			log.Fatal("can't get environment variabes ", err)
		}
		for k, v := range env {
			env := k + "=" + v
			envSlice = append(envSlice, env)
		}

		require.Equal(t, []string(nil), envSlice)
	})

	t.Run("one empty file", func(t *testing.T) {
		var envSlice []string
		inputPath := "testdata/testDir/"
		err := os.Mkdir(inputPath, 0777)
		file, err := os.Create(inputPath + "emptyFile")
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
		env, err := ReadDir(inputPath)
		if err != nil {
			log.Fatal("can't get environment variabes ", err)
		}
		for k, v := range env {
			env := k + "=" + v
			envSlice = append(envSlice, env)
		}

		require.Equal(t, []string{"emptyFile="}, envSlice)
	})
	t.Run("first line with several spaces", func(t *testing.T) {
		var envSlice []string
		inputPath := "testdata/testDir/"
		err := os.Mkdir(inputPath, 0777)
		file, err := os.Create(inputPath + "testFile")
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString(" a b  c d  ")
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
		env, err := ReadDir(inputPath)
		if err != nil {
			log.Fatal("can't get environment variabes ", err)
		}
		for k, v := range env {
			env := k + "=" + v
			envSlice = append(envSlice, env)
		}
		require.Equal(t, []string{"testFile= a b  c d"}, envSlice)
	})

	t.Run("several = in filename", func(t *testing.T) {
		var envSlice []string
		inputPath := "testdata/testDir/"
		err := os.Mkdir(inputPath, 0777)
		file, err := os.Create(inputPath + "==tes=tF=i=le===")
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString(" a b  c d  ")
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
		if err != nil {
			log.Fatal(err)
		}
		env, err := ReadDir(inputPath)
		if err != nil {
			log.Fatal("can't get environment variabes ", err)
		}
		for k, v := range env {
			env := k + "=" + v
			envSlice = append(envSlice, env)
		}
		require.Equal(t, []string{"testFile= a b  c d"}, envSlice)
	})
}
