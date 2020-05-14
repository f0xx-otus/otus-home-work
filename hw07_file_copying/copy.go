package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFromIsEmpty           = errors.New("source file is not set")
	ErrToIsEmpty             = errors.New("destination file is not set")
	ErrNegativeOffsetOrLimit = errors.New("offset and limit can't be negative")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrFromIsEmpty
	}
	if toPath == "" {
		return ErrToIsEmpty
	}
	if offset < 0 || limit < 0 {
		return ErrNegativeOffsetOrLimit
	}

	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0644)
	defer func() {
		err = fileFrom.Close()
		if err != nil {
			log.Panic("can't close output file", err)
		}
	}()
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
		return err
	}

	fileTo, err := os.Create(toPath)
	defer func() {
		err = fileTo.Close()
		if err != nil {
			log.Panic("can't close output file", err)
		}
	}()
	if err != nil {
		return err
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fileFrom)
	_, err = io.CopyN(fileTo, barReader, limit)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}
