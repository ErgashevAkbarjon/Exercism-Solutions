package thefarm

import "fmt"

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, n int) (float64, error) {

	amount, err := calc.FodderAmount(n)
	if err != nil {
		return 0, err
	}

	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (amount * factor) / float64(n), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}

	return DivideFood(calc, n)
}

type InvalidCowsError struct {
	NumberOfCows  int
	CustomMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", e.NumberOfCows, e.CustomMessage)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(n int) error {
	switch {
	case n < 0:
		return &InvalidCowsError{n, "there are no negative cows"}
	case n == 0:
		return &InvalidCowsError{n, "no cows don't need food"}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
