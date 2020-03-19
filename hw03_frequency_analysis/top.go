package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const whiteSpaceString = " "
const emptyString = ""
const characterToSkip = "-"
const specialSymbols = `[0-9$&+,:;=?@#|'<>.^*()%!]`
const whiteSpaceSymbol = `\s+`
const sizeToReturn = 10

func Top10(in string) []string {
	patternForSymbols := regexp.MustCompile(specialSymbols)
	inputWithoutSymbols := strings.Trim(patternForSymbols.ReplaceAllString(in, emptyString), whiteSpaceString)

	if len(inputWithoutSymbols) == 0 {
		return []string{}
	}

	patterToSpace := regexp.MustCompile(whiteSpaceSymbol)
	preparedInString := patterToSpace.ReplaceAllString(inputWithoutSymbols, whiteSpaceString)
	splitInString := strings.Split(preparedInString, whiteSpaceString)

	freqMap := map[string]int{}
	for _, val := range splitInString {
		if val == characterToSkip {
			continue
		}
		freqMap[strings.ToLower(val)]++
	}

	keys := make([]string, 0, len(freqMap))
	for key := range freqMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	resSize := len(keys)
	if resSize > sizeToReturn {
		resSize = sizeToReturn
	}

	return keys[:resSize]
}
