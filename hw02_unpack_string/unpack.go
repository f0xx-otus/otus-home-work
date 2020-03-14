package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var output string
	runeSlice := []rune(input)

	if input == "" {
		return "", nil
	}

	for i, r := range runeSlice {
		if unicode.IsLetter(r) {
			output += string(r)
			if i < (len(runeSlice) - 1) {
				if unicode.IsDigit(runeSlice[i+1]) {
					output += MultiplyLetters(runeSlice, i)
				}
			}
		}
		if unicode.IsDigit(r) {
			if i < (len(runeSlice) - 1) {
				if unicode.IsDigit(runeSlice[i+1]) {
					return "", ErrInvalidString
				}
			}
		}
	}
	return output, nil
}

func MultiplyLetters(runeSlice []rune, i int) string {
	letter := string(runeSlice[i])
	var multiLetter string
	for j := 1; j < int(runeSlice[i+1]-'0'); j++ {
		multiLetter += letter
	}
	return multiLetter
}
