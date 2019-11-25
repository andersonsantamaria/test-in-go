package operations

import (
	"ejercicios-de-prueba/src/operations"
	"ejercicios-de-prueba/test/builders"
	"testing"
)

var item builders.TestDataSum

func TestSumFunction(t *testing.T) {
	for _, item := range builders.ListDataSum() {
		result, error := operations.SumThreeNumbers(item.FirstNumber, item.SecondNumber, item.ThirdNumber)

		if item.HasError {
			if error == nil {
				t.Errorf("SumThreeNumbers() with args %v, %v, %v: FAILED, expected an error but got value '%v' ", item.FirstNumber, item.SecondNumber, item.ThirdNumber, result)
			} else {
				t.Logf("SumThreeNumbers() with args %v, %v, %v: PASSED, expected an error and got an error '%v' ", item.FirstNumber, item.SecondNumber, item.ThirdNumber, error)
			}
		} else {
			if result != item.Result {
				t.Errorf("SumThreeNumbers() with args %v, %v, %v: FAILED, expected %v but got value '%v' ", item.FirstNumber, item.SecondNumber, item.ThirdNumber, item.Result, result)
			} else {
				t.Logf("SumThreeNumbers() with args %v, %v, %v: PASSED, expected %v and got '%v' ", item.FirstNumber, item.SecondNumber, item.ThirdNumber, item.Result, result)
			}
		}
	}
}
