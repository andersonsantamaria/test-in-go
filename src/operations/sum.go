package operations

import "errors"

// SumTwoNumbers function add two int numbers
func SumTwoNumbers(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

// SumThreeNumbers function add three int numbers
func SumThreeNumbers(firstNumber, secondNumber, thirdNumber int) (result int, err error) {
	if secondNumber > 2 {
		return -1, errors.New("The second Argument cannot be greather than 2")
	}
	return firstNumber + secondNumber + thirdNumber, nil
}
