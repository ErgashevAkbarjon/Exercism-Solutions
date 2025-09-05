package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for _, v := range word {

		letterRepeats := unicode.IsLetter(v) &&
			strings.Count(word, string(v)) > 1

		if letterRepeats {
			return false
		}
	}

	return true
}
