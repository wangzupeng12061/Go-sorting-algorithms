package count

import (
	"reflect"
	"testing"
)

// TestCountingSort tests the CountingSort function
func TestCountingSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{4, 2, 2, 8, 3, 3, 1}, expected: []int{1, 2, 2, 3, 3, 4, 8}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 1, 1, 1, 1}, expected: []int{1, 1, 1, 1, 1}},
		{input: []int{}, expected: []int{}},
		{input: []int{10, -10, 0, 5, -5}, expected: []int{-10, -5, 0, 5, 10}},
	}

	for _, test := range tests {
		result := CountSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CountingSort(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
