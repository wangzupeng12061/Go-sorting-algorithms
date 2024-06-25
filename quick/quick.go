package quick

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	pivotIndex := randGen.Intn(len(arr))
	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]
	left, right := Partition(arr)
	QuickSort(arr[:left+1])
	QuickSort(arr[right:])
}
func Partition(arr []int) (int, int) {
	pivot := arr[len(arr)-1]
	left := -1
	right := len(arr)
	i := 0
	for i < right {
		if arr[i] < pivot {
			arr[i], arr[left+1] = arr[left+1], arr[i]
			left++
			i++
		} else if arr[i] == pivot {
			i++
		} else {
			arr[i], arr[right-1] = arr[right-1], arr[i]
			right--
		}
	}
	return left, right
}
