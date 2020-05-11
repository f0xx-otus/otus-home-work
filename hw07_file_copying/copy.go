package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFromIsEmpty           = errors.New("source file is not set")
	ErrToIsEmpty             = errors.New("destination file is not set")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrFromIsEmpty
	}

	if toPath == "" {
		return ErrToIsEmpty
	}

	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	fileInfo, err := fileFrom.Stat()
	if err != nil {
		return err
	}
	if limit == 0 {
		limit = fileInfo.Size()
	}
	if fileInfo.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	_, err = fileFrom.Seek(offset, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fileTo, err := os.Create(to)
	if err != nil {
		return err
	}
	_, err = io.CopyN(fileTo, fileFrom, limit)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Done")
		} else {
			fmt.Println(err)
			return err
		}
	}

	err = fileFrom.Close()
	if err != nil {
		return err
	}
	return nil
}
