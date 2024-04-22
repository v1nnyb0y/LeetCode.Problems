package d2

import (
	"fmt"
	"log"
	"math"
)

type Cafe struct {
	N      int
	prices []int
}

func (c *Cafe) inputFunc() {
	inputNumber := func() int {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
	c.N = inputNumber()
	for i := 0; i < c.N; i++ {
		c.prices = append(c.prices, inputNumber())
	}
}

func (c *Cafe) processing() {
	dp := make([][]int, c.N+1)
	for i := range dp {
		dp[i] = make([]int, c.N+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = 0
	for i := 1; i <= c.N; i++ {
		price := c.prices[i-1]
		for j := 0; j <= i; j++ {
			if j >= 1 {
				dp[i][j-1] = min(dp[i-1][j], dp[i][j-1])
			}
			if price > 100 && j+1 <= c.N {
				dp[i][j+1] = min(dp[i-1][j]+price, dp[i][j+1])
			} else {
				dp[i][j] = min(dp[i-1][j]+price, dp[i][j])
			}
		}
	}

	imin := 0
	minval := math.MaxInt32
	//for _, val := range dp {
	//	fmt.Println(val)
	//}
	for j := 0; j <= c.N; j++ {
		if minval >= dp[c.N][j] {
			minval = dp[c.N][j]
			imin = j
		}
	}
	// fmt.Println()

	spendCoupons := 0
	days := []int{}

	currJ := imin
	levelPrice := minval
	for i := c.N - 1; i > 0; i-- {
		price := c.prices[i]
		// fmt.Println(price, levelPrice, dp[i][currJ], price)
		switch {
		case currJ-1 >= 0 && price > 100 && dp[i][currJ-1]+price == levelPrice:
			{
				levelPrice -= price
				currJ -= 1
			}
		case dp[i][currJ]+price == levelPrice:
			{
				levelPrice -= price
			}
		case currJ+1 <= c.N && dp[i][currJ+1] == levelPrice:
			{
				days = append(days, i+1)
				currJ++
				spendCoupons++
			}
		}
	}

	// fmt.Println()

	fmt.Println(minval)
	fmt.Println(imin, spendCoupons)

	for i := len(days) - 1; i >= 0; i-- {
		fmt.Println(days[i])
	}
}

func CafeProcessing() {
	solution := &Cafe{}
	solution.inputFunc()
	solution.processing()
}
