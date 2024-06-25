package xor

// 给定一个数组，数组有多个数量是偶数的数，有两个数量是奇数的数，找到这两个数量是奇数的数
func Xor(arr []int) (int, int) {
	result := 0
	for _, number := range arr {
		result = result ^ number
	}
	rightOne := result & (^result + 1)
	result1 := 0
	for _, number := range arr {
		if rightOne&number == 0 {
			result1 = result ^ number
		}
	}
	return result1, result1 ^ result

}
