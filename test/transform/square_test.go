package transform

import (
	"ejercicios-de-prueba/src/transform"
	"reflect"
	"testing"
)

func TestTransformSquare(t *testing.T) {
	sliceToTransform := []int{1, 2, 3, 4, 5}
	expectedSlice := []int{1, 4, 9, 16, 25}
	result := transform.SquareSlice(sliceToTransform)

	if reflect.DeepEqual(expectedSlice, result) {
		t.Log("SquareSlice PASSED")
	} else {
		t.Errorf("SquareSlice FAILED, expected %v but got %v", expectedSlice, result)
	}
}
