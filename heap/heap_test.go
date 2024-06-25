package heap

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{12, 11, 13, 5, 6, 7}, expected: []int{5, 6, 7, 11, 12, 13}},
		{input: []int{64, 34, 25, 12, 22, 11, 90}, expected: []int{11, 12, 22, 25, 34, 64, 90}},
		{input: []int{5, 1, 4, 2, 8}, expected: []int{1, 2, 4, 5, 8}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{2, 1}, expected: []int{1, 2}},
		{input: []int{3, 3, 3}, expected: []int{3, 3, 3}},
	}

	for _, test := range tests {
		result := make([]int, len(test.input))
		copy(result, test.input)
		HeapSortCircle(result)
		HeapSort(result)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("HeapSort(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
