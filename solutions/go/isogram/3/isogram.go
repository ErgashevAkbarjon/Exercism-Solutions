package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {

	for _, v := range strings.ToLower(word) {

		letterRepeats := unicode.IsLetter(v) &&
			strings.Count(word, string(v)) > 1

		if letterRepeats {
			return false
		}
	}

	return true
}
