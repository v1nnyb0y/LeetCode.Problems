package contest

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type ClosestStop struct {
	N, K             int
	Stops, TestCases []int
}

func (cs *ClosestStop) inputFunc() {
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
	cs.N, cs.K = inputNumber(), inputNumber()
	cs.Stops = inputArr(cs.N)
	cs.TestCases = inputArr(cs.K)
}

func (cs *ClosestStop) outputFunc() {
	for _, x := range cs.TestCases {
		index := sort.Search(cs.N, func(i int) bool { return cs.Stops[i] >= x })
		var closestIndex int
		if index == 0 {
			closestIndex = 0
		} else if index == cs.N {
			closestIndex = cs.N - 1
		} else {
			leftDist := x - cs.Stops[index-1]
			rightDist := cs.Stops[index] - x

			if leftDist < rightDist || (leftDist == rightDist && index > 0) {
				closestIndex = index - 1
			} else {
				closestIndex = index
			}
		}

		// Output the 1-based index of the closest bus stop
		fmt.Println(closestIndex + 1)
	}
}

func ClosestStopProcessing() {
	solution := &ClosestStop{}
	solution.inputFunc()
	solution.outputFunc()
}
