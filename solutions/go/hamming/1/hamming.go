package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	first, second := []byte(a), []byte(b)
	diffs := 0

	if len(first) != len(second) {
		return diffs, fmt.Errorf("a and b should have similar length")
	}

	for i, v := range first {
		if v != second[i] {
			diffs++
		}
	}

	return diffs, nil
}
