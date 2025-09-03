package scrabble

import (
	"strings"
)

func Score(word string) int {
	scoreTable := map[int]string{
		1:  "A, E, I, O, U, L, N, R, S, T",
		2:  "D, G",
		3:  "B, C, M, P",
		4:  "F, H, V, W, Y",
		5:  "K",
		8:  "J, X",
		10: "Q, Z",
	}

	score := 0
	word = strings.ToUpper(word)

	for _, r := range word {
		for s, v := range scoreTable {
			if strings.ContainsRune(v, r) {
				score += s
			}
		}
	}

	return score
}
