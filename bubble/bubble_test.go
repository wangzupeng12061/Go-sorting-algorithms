package bubble

import (
	"reflect"
	"testing"
)

// TestBubbleSort 测试冒泡排序函数
func TestBubbleSort(t *testing.T) {
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
		BubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("BubbleSort(%v) = %v; expected %v", test.input, test.input, test.expected)
		}
	}
}
