package solutions

import "sort"

func threeSum(nums []int) [][]int {
	triplets := [][]int{}

	tripmap := make(map[[3]int]bool)

	sort.Ints(nums)

	if len(nums) <= 2 {
		return triplets
	}
	if nums[0] > 0 {
		return triplets
	}

	for num := range nums[:len(nums)-2] {
		left := num + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[num] + nums[left] + nums[right]
			if currentSum == 0 {
				triplet := [3]int{nums[num], nums[left], nums[right]}
				tripmap[triplet] = true
				left++
				right--
			} else if currentSum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	for k := range tripmap {
		tempArray := []int{}
		for _, val := range k {
			tempArray = append(tempArray, val)
		}
		triplets = append(triplets, tempArray)
	}

	return triplets
}
