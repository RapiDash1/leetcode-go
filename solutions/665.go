package solutions

func checkPossibility(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	diffCount := 0

	for i := range nums[:len(nums)-1] {
		if nums[len(nums)-i-2] > nums[len(nums)-i-1] {
			diffCount++

			if diffCount > 1 {
				return false
			}

			if i != 0 && nums[len(nums)-i] <= nums[len(nums)-i-2] {
				nums[len(nums)-i-2] = nums[len(nums)-i-1]
			}
		}
	}
	return true
}
