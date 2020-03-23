package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const (
	whiteSpaceString = " "
	emptyString      = ""
	characterToSkip  = "-"
	specialSymbols   = `[0-9$&+,:;=?@#|'<>.^*()%!]`
	whiteSpaceSymbol = `\s+`
	sizeToReturn     = 10
)

var (
	patternForSymbols = regexp.MustCompile(specialSymbols)
	patternToSpace    = regexp.MustCompile(whiteSpaceSymbol)
)

func Top10(in string) []string {
	inputWithoutSymbols := strings.Trim(patternForSymbols.ReplaceAllString(in, emptyString), whiteSpaceString)

	if len(inputWithoutSymbols) == 0 {
		return nil
	}

	preparedInString := patternToSpace.ReplaceAllString(inputWithoutSymbols, whiteSpaceString)
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
