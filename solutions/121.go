package solutions

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minVal := prices[0]
	maxProfit := 0

	for _, num := range prices {
		if num < minVal {
			minVal = num
		}
		if num-minVal > maxProfit {
			maxProfit = num - minVal
		}
	}
	return maxProfit
}
