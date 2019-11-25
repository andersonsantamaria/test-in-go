package builders

// TestDataSum is a struct to test the operations.Sum function
type TestDataSum struct {
	FirstNumber  int
	SecondNumber int
	ThirdNumber  int
	Result       int
	HasError     bool
}

// ListDataSum is a list of data for the operations.Sum function
func ListDataSum() []TestDataSum {
	return []TestDataSum{
		{1, 1, 0, 2, false},
		{1, 2, 1, 4, false},
		{1, 2, 2, 5, false},
		{1, 2, 3, 6, false},
		{1, 2, 4, 7, false},
		{1, 2, 5, 8, false},
		{1, 3, 6, 9, true}}
}
