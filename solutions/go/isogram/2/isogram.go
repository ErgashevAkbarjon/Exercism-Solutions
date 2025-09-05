package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {

	for _, v := range strings.ToLower(word) {

		if !unicode.IsLetter(v) {
			continue
		}

		if strings.Count(word, string(v)) > 1 {
			return false
		}
	}

	return true
}
