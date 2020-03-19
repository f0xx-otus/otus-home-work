package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var output string
	runeSlice := []rune(input)
	var isNextSymbolEscaped bool
	if input == "" {
		return "", nil
	}
	for i, r := range runeSlice {
		if r == '\\' && len(runeSlice) == 1 {
			return "", ErrInvalidString
		}
		if r == '\\' && !isNextSymbolEscaped {
			isNextSymbolEscaped = true
			continue
		}
		if unicode.IsDigit(r) && !isNextSymbolEscaped {
			err := twoDigitCheck(i, runeSlice)
			if err != nil {
				return "", err
			}
		}
		if unicode.IsLetter(r) && i > 0 && isNextSymbolEscaped {
			isNextSymbolEscaped = false
			continue
		}
		if unicode.IsLetter(r) || isNextSymbolEscaped {
			output += string(r)
			isNextSymbolEscaped = false
			if i < (len(runeSlice) - 1) {
				if unicode.IsDigit(runeSlice[i+1]) {
					output += multiplyLetters(r, runeSlice[i+1])
				}
			}
		}
	}
	if len(runeSlice) > 0 && output == "" {
		return "", ErrInvalidString
	}
	return output, nil
}

func multiplyLetters(rune rune, i rune) string {
	return strings.Repeat(string(rune), int(i-'0')-1)
}

func twoDigitCheck(i int, runeSlice []rune) error {
	if i < (len(runeSlice) - 1) {
		if unicode.IsDigit(runeSlice[i+1]) {
			return ErrInvalidString
		}
	}
	return nil
}
