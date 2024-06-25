package swap

import (
	"reflect"
	"testing"
)

// TestSwapSort 测试选择排序算法
func TestSwapSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{64, 34, 25, 12, 22, 11, 90}, expected: []int{11, 12, 22, 25, 34, 64, 90}},
		{input: []int{5, 1, 4, 2, 8}, expected: []int{1, 2, 4, 5, 8}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{2, 1}, expected: []int{1, 2}},
		{input: []int{3, 3, 3}, expected: []int{3, 3, 3}},
	}

	for _, test := range tests {
		result := SwapSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SwapSort(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
