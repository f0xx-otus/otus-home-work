package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("directory without files", func(t *testing.T) {
		var envSlice []string
		inputPath := "testdata/emptyDir"
		err := os.Mkdir(inputPath, 0777)
		require.NoError(t, err)
		defer func() {
			err = os.Remove(inputPath)
			require.NoError(t, err)
		}()
		env, err := ReadDir(inputPath)
		require.NoError(t, err)
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
		require.NoError(t, err)
		file, err := os.Create(inputPath + "emptyFile")
		require.NoError(t, err)
		err = file.Close()
		require.NoError(t, err)
		defer func() {
			err = os.RemoveAll(inputPath)
			require.NoError(t, err)
		}()
		env, err := ReadDir(inputPath)
		require.NoError(t, err)
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
		require.NoError(t, err)
		file, err := os.Create(inputPath + "testFile")
		require.NoError(t, err)
		_, err = file.WriteString(" a b  c d  ")
		require.NoError(t, err)
		err = file.Close()
		require.NoError(t, err)
		defer func() {
			err = os.RemoveAll(inputPath)
			require.NoError(t, err)
		}()
		env, err := ReadDir(inputPath)
		require.NoError(t, err)
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
		require.NoError(t, err)
		file, err := os.Create(inputPath + "==tes=tF=i=le===")
		require.NoError(t, err)
		_, err = file.WriteString(" a b  c d  ")
		require.NoError(t, err)
		err = file.Close()
		require.NoError(t, err)
		defer func() {
			err = os.RemoveAll(inputPath)
			require.NoError(t, err)
		}()
		require.NoError(t, err)
		env, err := ReadDir(inputPath)
		require.NoError(t, err)
		for k, v := range env {
			env := k + "=" + v
			envSlice = append(envSlice, env)
		}
		require.Equal(t, []string{"testFile= a b  c d"}, envSlice)
	})
}
