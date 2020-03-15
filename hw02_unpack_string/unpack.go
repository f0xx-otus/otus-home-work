package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	fmt.Println(input)
	var output string
	runeSlice := []rune(input)
	isNextSymbolEscaped := false

	if input == "" {
		return "", nil
	}

	for i, r := range runeSlice {
		if r == '\\' && isNextSymbolEscaped == false {
			isNextSymbolEscaped = true
			continue
		}
		if unicode.IsDigit(r) && isNextSymbolEscaped == false {
			if i < (len(runeSlice) - 1) {
				if unicode.IsDigit(runeSlice[i+1]) {
					return "", ErrInvalidString
				}
			}
		}
		if unicode.IsLetter(r) || isNextSymbolEscaped {
			output += string(r)
			isNextSymbolEscaped = false
			if i < (len(runeSlice) - 1) {
				if unicode.IsDigit(runeSlice[i+1]) {
					output += MultiplyLetters(runeSlice, i)
				}
			}
		}
		if len(runeSlice) > 0 && output == "" {
			return "", ErrInvalidString
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
