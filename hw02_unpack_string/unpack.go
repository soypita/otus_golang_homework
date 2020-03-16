package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const EscapeSymbolRune = '\\'
const EscapeSymbolString = `\`
const stringPart = 2
const prevOffset = 1

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	if _, err := strconv.Atoi(s); err == nil {
		return "", ErrInvalidString
	}

	var builder strings.Builder
	inputSymbols := []rune(s)
	countOfScreens := 0

	for i, val := range inputSymbols {
		if val == EscapeSymbolRune {
			countOfScreens++
			continue
		}

		if unicode.IsDigit(val) {
			if i == 0 || (unicode.IsDigit(inputSymbols[i-1]) && inputSymbols[i-2] != EscapeSymbolRune) {
				return "", ErrInvalidString
			}

			builder.WriteString(strings.Repeat(EscapeSymbolString, countOfScreens/stringPart))

			if countOfScreens%2 != 0 {
				builder.WriteString(string(val))
			} else {
				countToRepeat, err := strconv.Atoi(string(val))
				if err != nil {
					return "", ErrInvalidString
				}
				builder.WriteString(strings.Repeat(string(inputSymbols[i-1]), countToRepeat-prevOffset))
			}
			countOfScreens = 0
		} else if unicode.IsLetter(val) {
			builder.WriteString(strings.Repeat(EscapeSymbolString, countOfScreens/stringPart))
			builder.WriteString(string(val))
			countOfScreens = 0
		}
	}

	if countOfScreens != 0 {
		if countOfScreens%2 != 0 {
			return "", ErrInvalidString
		}
		builder.WriteString(strings.Repeat(EscapeSymbolString, countOfScreens/stringPart))
	}

	return builder.String(), nil
}
