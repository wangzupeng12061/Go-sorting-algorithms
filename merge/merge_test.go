package merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, test := range tests {
		result := MergeSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MergeSort(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
