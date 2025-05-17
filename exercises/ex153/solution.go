package ex153

func findMin(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] > nums[1] {
			return nums[1]
		}
		return nums[0]
	}

	middle := len(nums) / 2
	left, right := nums[:middle], nums[middle:]

	leftMin := findMin([]int{left[0], left[len(left)-1]})
	rightMin := findMin([]int{right[0], right[len(right)-1]})

	if leftMin > rightMin {
		return findMin(right)
	}
	return findMin(left)
}
