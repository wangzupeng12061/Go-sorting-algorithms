package swap

func SwapSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		//minVal := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] { //if arr[j]<minVal
				//minVal = arr[j]
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}
	return arr

}
