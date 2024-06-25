package heap

func HeapInsert(arr []int, index int) {

	for arr[index] > arr[(index-1)/2] {
		par := (index - 1) / 2
		arr[index], arr[par] = arr[par], arr[index]
		index = par
	}

}

func Heapify(arr []int, index int, heapSize int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index
	if left < heapSize && arr[left] > arr[index] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		Heapify(arr, largest, heapSize)
	}
}
func HeapifyCircle(arr []int, index int, heapSize int) {

	for 2*index+1 < heapSize {
		left := 2*index + 1
		right := 2*index + 2
		largest := left
		if right < heapSize && arr[right] > arr[left] {
			largest = right
		}
		if arr[largest] < arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
	}
}
func HeapSort(arr []int) {
	heapSize := len(arr)
	for i := heapSize/2 - 1; i >= 0; i-- {
		Heapify(arr, i, heapSize)
	}
	for i := heapSize - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Heapify(arr, 0, i)
	}
}
func HeapSortCircle(arr []int) {
	heapSize := len(arr)
	for i := 0; i < heapSize; i++ {
		HeapInsert(arr, i)
	}

	for i := heapSize - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapifyCircle(arr, 0, i)
	}
}
