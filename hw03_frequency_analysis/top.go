package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const splitChar = " "
const characterToSkip = "-"
const specialSymbols = `[0-9$&+,:;=?@#|'<>.^*()%!]`
const whiteSpaceSymbol = `\s+`
const sizeToReturn = 10

func Top10(in string) []string {
	patternForSymbols := regexp.MustCompile(specialSymbols)
	inputWithoutSymbols := strings.Trim(patternForSymbols.ReplaceAllString(in, ""), " ")

	if len(inputWithoutSymbols) == 0 {
		return []string{}
	}

	space := regexp.MustCompile(whiteSpaceSymbol)
	preparedInString := space.ReplaceAllString(inputWithoutSymbols, " ")
	splitInString := strings.Split(preparedInString, splitChar)

	freqMap := map[string]int{}
	for _, val := range splitInString {
		if val == characterToSkip {
			continue
		}
		freqMap[strings.ToLower(val)]++
	}

	it := 0
	keys := make([]string, len(freqMap))
	for key := range freqMap {
		keys[it] = key
		it++
	}

	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	resSize := len(keys)
	if resSize > sizeToReturn {
		resSize = sizeToReturn
	}

	res := make([]string, resSize)
	for i, val := range keys {
		if i == resSize {
			break
		}
		res[i] = val
	}
	return res
}
