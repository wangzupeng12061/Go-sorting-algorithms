package quick

func sortColors(nums []int) {
	std := 1
	i := 0
	left := -1
	right := len(nums)
	for i < right {
		if nums[i] < std {
			nums[i], nums[left+1] = nums[left+1], nums[i]
			i++
			left++
		} else if nums[i] == std {
			i++
		} else {
			nums[i], nums[right-1] = nums[right-1], nums[i]
			right--
		}
	}
}
