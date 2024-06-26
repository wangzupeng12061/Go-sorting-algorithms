package radix

import (
	"reflect"
	"testing"
)

// TestRadixSort tests the RadixSort function
func TestRadixSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{170, 45, 75, 90, 802, 24, 2, 66}, expected: []int{2, 24, 45, 66, 75, 90, 170, 802}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 1, 1, 1, 1}, expected: []int{1, 1, 1, 1, 1}},
		{input: []int{}, expected: []int{}},
		{input: []int{10, -10, 0, 5, -5}, expected: []int{-10, -5, 0, 5, 10}}, // RadixSort assumes non-negative integers, so modify input accordingly.
	}

	for _, test := range tests {
		result := RadixSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("RadixSort(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
