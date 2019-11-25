package operations

import (
	"testing"
)

func TestSumPositiveNumbers(t *testing.T) {
	firstNumber := 15
	secondNumber := 16
	expected := 31
	result := SumTwoNumbers(firstNumber, secondNumber)

	if expected != result {
		t.Errorf("The sum operation failed, expected %v, got %v", expected, result)
	}
}

func TestSumNegativeNumbers(t *testing.T) {
	firstNumber := -15
	secondNumber := -16
	expected := -31
	result := SumTwoNumbers(firstNumber, secondNumber)

	if expected != result {
		t.Errorf("The sum operation failed, expected %v, got %v", expected, result)
	}
}
