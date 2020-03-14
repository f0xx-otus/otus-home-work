package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	fmt.Println(input)
	var output string
	if output == "" {
		return "invalid string", ErrInvalidString
	}

	return output, nil
}
