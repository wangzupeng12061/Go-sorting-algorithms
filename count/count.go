package count

func CountSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	ma, mi := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > ma {
			ma = arr[i]
		}
		if arr[i] < mi {
			mi = arr[i]
		}

	}
	count := make([]int, ma-mi+1)
	for _, num := range arr {
		count[num-mi]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[count[arr[i]-mi]-1] = arr[i]
		count[arr[i]-mi]--
	}
	return result
}
