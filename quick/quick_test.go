package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)
		QuickSort(inputCopy)
		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("QuickSort(%v) = %v; expected %v", test.input, inputCopy, test.expected)
		}
	}
}
