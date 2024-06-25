package quick

import (
	"reflect"
	"testing"
)

// TestSortColors tests the sortColors function
func TestSortColors(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{2, 0, 1}, expected: []int{0, 1, 2}},
		{input: []int{0}, expected: []int{0}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{2, 2, 2, 1, 1, 0, 0, 0}, expected: []int{0, 0, 0, 1, 1, 2, 2, 2}},
	}

	for _, test := range tests {
		sortColors(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("sortColors(%v) = %v; expected %v", test.input, test.input, test.expected)
		}
	}
}
