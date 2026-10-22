func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	profit := 0
	l, r := 0, 1

	for r < len(prices) {
		if prices[l] >= prices[r] {
			l = r
			r = l + 1
		} else {
			if prices[r] - prices[l] > profit {
				profit = prices[r] - prices[l]
			}
			r++
		}
	}
	
	return profit
}
