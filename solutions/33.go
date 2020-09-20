package solutions

func pivot(nums []int) int {
	low := 0
	high := len(nums) - 1
	mid := 0

	for low < high {
		if mid == (low+high)/2 {
			break
		}
		mid = (low + high) / 2
		if mid == low {
			return low
		}
		if nums[mid] < nums[mid-1] {
			return mid
		}

		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid
		}
	}

	return low
}

func binarySearch(nums []int, low, high, target int) int {
	if target == nums[low] {
		return low
	} else if target == nums[high] {
		return high
	}

	mid := 0
	for low < high {
		if mid == (low+high)/2 {
			break
		}
		mid = (low + high) / 2
		if mid == low {
			break
		}
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid
		} else if nums[mid] > target {
			high = mid
		}

	}

	return -1
}

func search(nums []int, target int) int {
	pivotIndex := pivot(nums)
	if target == nums[pivotIndex] {
		return pivotIndex
	}
	if nums[0] <= nums[len(nums)-1] {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	if target < nums[len(nums)-1] {
		return binarySearch(nums, pivotIndex, len(nums)-1, target)
	} else if target > nums[len(nums)-1] {
		return binarySearch(nums, 0, pivotIndex, target)
	}
	return len(nums) - 1
}
