package merge

/*
给你一个整数数组 nums ，按要求返回一个新数组 counts 。
数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
*/

type Key struct {
	Index int
	Val   int
}

var fan []int
var re map[Key]int

func CountSmaller(nums []int) []int {
	for index, value := range nums {
		re[Key{Index: index, Val: value}] = 0
	}
	re = make(map[Key]int)
	_ = mergeSort(nums)
	for _, v := range re {
		fan = append(fan, v)
	}
	return fan
}
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)

}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < len(left) {
		result = append(result, left[i:]...)
	} else {
		result = append(result, right[j:]...)
	}
	return result
}
