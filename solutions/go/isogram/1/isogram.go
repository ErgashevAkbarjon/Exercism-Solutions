package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ToLower(word)

	for _, v := range word {
		count := strings.Count(word, string(v))

		if count > 1 {
			return false
		}
	}

	return true
}
