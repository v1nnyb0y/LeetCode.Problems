package dp

import "math"

func maxProfit(prices []int) int {

	cur_hold, cur_not_hold := math.MinInt64, 0

	for _, stock_price := range prices {

		prev_hold, prev_not_hold := cur_hold, cur_not_hold

		// either keep hold, or buy in stock today at stock price
		cur_hold = max(prev_hold, prev_not_hold-stock_price)

		// either keep not-hold, or sell out stock today at stock price
		cur_not_hold = max(prev_not_hold, prev_hold+stock_price)

	}

	return cur_not_hold

}
