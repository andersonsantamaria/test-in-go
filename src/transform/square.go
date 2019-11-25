package transform

// SquareSlice is a function to squre each item of a slice
func SquareSlice(sliceToTransform []int) []int {
	squareSlice := make([]int, len(sliceToTransform))

	for index, value := range sliceToTransform {
		squareSlice[index] = value * value
	}

	return squareSlice
}
