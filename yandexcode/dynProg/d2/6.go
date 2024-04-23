package d2

import (
	"fmt"
	"log"
)

type LCS struct {
	N, M       int
	NARR, MARR []int
}

func (lcs *LCS) inputFunc() {
	inputNumber := func() int {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
	inputArr := func(lenArr int) []int {
		var arr = make([]int, lenArr)
		for i := 0; i < lenArr; i++ {
			_, err := fmt.Scan(&arr[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		return arr
	}
	lcs.N = inputNumber()
	lcs.NARR = inputArr(lcs.N)
	lcs.M = inputNumber()
	lcs.MARR = inputArr(lcs.M)
}

func (lcs *LCS) process() {
	dp := make([][]int, lcs.N+1)
	for i := range dp {
		dp[i] = make([]int, lcs.M+1)
	}

	for i := 1; i <= lcs.N; i++ {
		for j := 1; j <= lcs.M; j++ {
			if lcs.NARR[i-1] == lcs.MARR[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	i, j := lcs.N, lcs.M
	//for _, val := range dp {
	//	fmt.Println(val)
	//}
	//fmt.Println(dp[i][j])
	result := make([]int, 0, dp[i][j])
	for dp[i][j] > 0 {
		switch {
		case lcs.NARR[i-1] == lcs.MARR[j-1]:
			{
				result = append(result, lcs.NARR[i-1])
				i--
				j--
			}
		case dp[i-1][j] > dp[i][j-1]:
			{
				// fmt.Print(lcs.NARR[i-1], " ")
				i--
			}
		default:
			{
				// fmt.Print(lcs.MARR[j-1], " ")
				j--
			}
		}
	}

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i], " ")
	}
}

func LCSProcess() {
	solution := &LCS{}
	solution.inputFunc()
	solution.process()
}
