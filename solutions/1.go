package solutions

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	satisfiedIndices := make([]int, 2)
	for i, num := range nums {
		if _, ok := cache[target-num]; ok {
			satisfiedIndices[0] = i
			satisfiedIndices[1] = cache[target-num]
			return satisfiedIndices
		}
		cache[num] = i
	}
	return satisfiedIndices
}
