package solutions

func productExceptSelf(nums []int) []int {
	outputArray := make([]int, len(nums))

	if len(nums) == 0 {
		return outputArray
	}

	prod := 1
	for i, num := range nums {
		outputArray[i] = prod
		prod *= num
	}

	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		oldVal := nums[i]
		nums[i] = prod
		prod *= oldVal
	}

	for i := range nums {
		outputArray[i] *= nums[i]
	}

	return outputArray
}
