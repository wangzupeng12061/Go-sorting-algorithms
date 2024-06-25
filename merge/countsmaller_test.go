package merge

import (
	"reflect"
	"testing"
)

// TestCountSmallerElements tests the CountSmallerElements function
func TestCountSmallerElements(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 6, 1}, expected: []int{2, 1, 1, 0}},
		{input: []int{1, 2, 3, 4}, expected: []int{0, 0, 0, 0}},
		{input: []int{4, 3, 2, 1}, expected: []int{3, 2, 1, 0}},
		{input: []int{1, 1, 1, 1}, expected: []int{0, 0, 0, 0}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		result := CountSmaller(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CountSmallerElements(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
