package radix

func getMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ma := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > ma {
			ma = arr[i]
		}
	}
	return ma

}
func getMin(arr []int) int {
	mi := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < mi {
			mi = arr[i]
		}
	}
	return mi
}
func countingSort(arr []int, exp int) []int {
	n := len(arr)
	out := make([]int, n)
	count := make([]int, 10)
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		out[count[index]-1] = arr[i]
		count[index]--
	}
	return out
}
func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	ma := getMax(arr)
	mi := getMin(arr)
	if mi < 0 {
		for i := 0; i < len(arr); i++ {
			arr[i] -= mi
		}
	}
	for exp := 1; ma/exp > 0; exp *= 10 {
		arr = countingSort(arr, exp)
	}
	if mi < 0 {
		for i := 0; i < len(arr); i++ {
			arr[i] += mi
		}
	}
	return arr
}
