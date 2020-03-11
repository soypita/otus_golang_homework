package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const EscapeSymbol = `\`

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
		if string(val) == EscapeSymbol {
			countOfScreens++
			continue
		}
		if unicode.IsDigit(val) {
			if i == 0 || (unicode.IsDigit(inputSymbols[i-1]) && string(inputSymbols[i-2]) != EscapeSymbol) {
				return "", ErrInvalidString
			}

			builder.WriteString(strings.Repeat(EscapeSymbol, countOfScreens/2))

			if countOfScreens%2 != 0 {
				builder.WriteString(string(val))
			} else {
				countToRepeat, err := strconv.Atoi(string(val))
				if err != nil {
					return "", ErrInvalidString
				}
				builder.WriteString(strings.Repeat(string(inputSymbols[i-1]), countToRepeat-1))
			}
			countOfScreens = 0
		} else if unicode.IsLetter(val) {
			builder.WriteString(strings.Repeat(EscapeSymbol, countOfScreens/2))
			builder.WriteString(string(val))
			countOfScreens = 0
		}
	}

	if countOfScreens != 0 {
		if countOfScreens%2 != 0 {
			return "", ErrInvalidString
		}
		builder.WriteString(strings.Repeat(EscapeSymbol, countOfScreens/2))
	}

	return builder.String(), nil
}
