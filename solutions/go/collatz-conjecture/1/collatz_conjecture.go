package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n <= 0 {
		return 0, fmt.Errorf("n must be more than 0")
	}

	if n == 1 {
		return steps, nil
	}

	for {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		steps++

		if n == 1 {
			break
		}
	}

	return steps, nil
}
