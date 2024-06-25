package insert

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i
		for j > 0 && arr[j-1] > key {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = key
	}
	return arr

}
