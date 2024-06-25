package xor

import (
	"reflect"
	"sort"
	"testing"
)

// TestFindOddOccurrences 测试 findOddOccurrences 函数
func TestFindOddOccurrences(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{4, 2, 4, 5, 2, 3, 3, 1}, expected: []int{1, 5}},
		{input: []int{10, 3, 10, 3, 10, 5, 5, 7, 7, 9}, expected: []int{9, 10}},
		{input: []int{1, 1, 2, 2, 3, 4}, expected: []int{3, 4}},
	}

	for _, test := range tests {
		num1, num2 := Xor(test.input)
		result := []int{num1, num2}
		sort.Ints(result)
		sort.Ints(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("findOddOccurrences(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
