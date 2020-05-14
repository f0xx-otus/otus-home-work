package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset is greater than file size", func(t *testing.T) {
		inputPath := "testdata/input.txt"
		outputPath := "/tmp/out.txt"
		fileFrom, err := os.OpenFile(inputPath, os.O_RDONLY, 0644)
		if err != nil {
			log.Panicf("can't open input file")
		}
		offset, err := fileFrom.Stat()

		if err != nil {
			log.Panicf("can't get file properties")
		}
		err = Copy(inputPath, outputPath, offset.Size()+1, 0)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("negative offset", func(t *testing.T) {
		inputPath := "testdata/input.txt"
		outputPath := "/tmp/out.txt"

		err := Copy(inputPath, outputPath, -1, 0)
		require.Equal(t, ErrNegativeOffsetOrLimit, err)
	})

	t.Run("negative limit", func(t *testing.T) {
		inputPath := "testdata/input.txt"
		outputPath := "/tmp/out.txt"

		err := Copy(inputPath, outputPath, 0, -1)
		require.Equal(t, ErrNegativeOffsetOrLimit, err)
	})

	t.Run("from path is empty", func(t *testing.T) {
		inputPath := ""
		outputPath := "/tmp/out.txt"

		err := Copy(inputPath, outputPath, 0, 0)
		require.Equal(t, ErrFromIsEmpty, err)
	})

	t.Run("to path is empty", func(t *testing.T) {
		inputPath := "/tmp/in.txt"
		outputPath := ""

		err := Copy(inputPath, outputPath, 0, 0)
		require.Equal(t, ErrToIsEmpty, err)
	})

	t.Run("unsupported file", func(t *testing.T) {
		inputPath := "/dev/random"
		outputPath := "/tmp/out"

		err := Copy(inputPath, outputPath, 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("offset 2, limit 5", func(t *testing.T) {
		var inputPath = "testdata/input.txt"

		outFile, err := ioutil.TempFile("/tmp", "prefix")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err = os.Remove(outFile.Name())
			if err != nil {
				log.Panic("can't close output file", err)
			}
		}()

		err = Copy(inputPath, outFile.Name(), 3, 9)
		require.NoError(t, err)

		resultFile, err := os.OpenFile(outFile.Name(), os.O_RDONLY, 0644)
		if err != nil {
			log.Panicf("can't open input file")
		}
		output, err := ioutil.ReadAll(resultFile)
		require.NoError(t, err)

		require.Equal(t, "Documents", string(output))
	})
}
