package dp

import "math"

func MaxProfit(prices []int) int {
	maxProfit := math.MinInt32
	timeToBuy := prices[0]
	for _, price := range prices[1:] {
		if price < timeToBuy {
			timeToBuy = price
			continue
		}
		maxProfit = max(price-timeToBuy, maxProfit)
	}

	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}
