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
		if os.IsNotExist(err) {
			return errors.New("fileFrom is not exist")
		}
		return errors.New("can't open fileFrom")
	}

	fileInfo, err := fileFrom.Stat()
	if err != nil {
		return errors.New("can't get fileFrom properties")
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
		return errors.New("can't change offset")
	}

	fileTo, err := os.Create(to)
	if err != nil {
		return errors.New("can't create a file")
	}
	_, err = io.CopyN(fileTo, fileFrom, limit)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Done")
		} else {
			fmt.Println(err)
			return errors.New("can't copy file")
		}
	}

	err = fileFrom.Close()
	if err != nil {
		fmt.Println(err)
		return errors.New("can't close fileFrom")
	}
	return nil
}
