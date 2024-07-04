package contest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type MaxSquare struct {
	N, M int
	Box  [][]int
}

func (ms *MaxSquare) inputFunc() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	inputNumber := func() int {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		return n
	}
	inputArr := func(lenArr int) []int {
		a := make([]int, lenArr)
		for i := 0; i < lenArr; i++ {
			scanner.Scan()
			a[i], _ = strconv.Atoi(scanner.Text())
		}
		return a
	}

	ms.N, ms.M = inputNumber(), inputNumber()
	for i := 0; i < ms.N; i++ {
		ms.Box = append(ms.Box, inputArr(ms.M))
	}
}

func (ms *MaxSquare) outputFunc() {
	dp := make([][]int, ms.N+1)
	for i := 0; i < ms.N+1; i++ {
		dp[i] = make([]int, ms.M+1)
	}

	var maxSize, maxI, maxJ int
	for i := 1; i < ms.N+1; i++ {
		for j := 1; j < ms.M+1; j++ {
			if ms.Box[i-1][j-1] == 1 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSize {
					maxSize = dp[i][j]
					maxI = i - 1
					maxJ = j - 1
				}
			}
		}
	}

	fmt.Println(maxSize)
	fmt.Println(fmt.Sprintf("%d %d", maxI-maxSize+2, maxJ-maxSize+2))
}

func MaxSquareProcessing() {
	solution := &MaxSquare{}
	solution.inputFunc()
	solution.outputFunc()
}
