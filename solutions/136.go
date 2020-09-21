package solutions

func sum(nums *[]int) int {
	total := 0
	for _, num := range *nums {
		total += num
	}
	return total
}

func setSum(nums *[]int) int {
	setMap := make(map[int]bool)
	total := 0
	for _, num := range *nums {
		if _, ok := setMap[num]; !ok {
			setMap[num] = true
			total += num
		}
	}
	return total
}

func singleNumber(nums []int) int {
	return 2*setSum(&nums) - sum(&nums)
}
