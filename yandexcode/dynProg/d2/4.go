package d2

import (
	"fmt"
	"log"
)

type Step struct {
	incI, incJ int
}

type Pos struct {
	I, J int
}
type HorseCount struct {
	N, M int
}

func (hc *HorseCount) inputFunc() {
	readArray := func(lenArr int) []int {
		var arr = make([]int, lenArr)
		for i := 0; i < len(arr); i++ {
			_, err := fmt.Scan(&arr[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		return arr
	}
	sizes := readArray(2)
	hc.N, hc.M = sizes[0], sizes[1]
}

func HorseCountProcess() {
	solution := &HorseCount{}
	solution.inputFunc()
	steps := []Step{
		{incI: 2, incJ: 1},
		{incI: 1, incJ: 2},
	}

	dp := make([][]int, solution.N)
	for i := 0; i < solution.N; i++ {
		dp[i] = make([]int, solution.M)
	}
	dp[0][0] = 1

	queue := []Pos{{I: 0, J: 0}}
	visited := make(map[[2]int]struct{})
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		key := [2]int{cur.I, cur.J}

		for _, step := range steps {
			newCur := Pos{I: cur.I + step.incI, J: cur.J + step.incJ}
			if newCur.I >= solution.N {
				continue
			}
			if newCur.J >= solution.M {
				continue
			}

			if _, ok := visited[key]; !ok {
				dp[newCur.I][newCur.J] += dp[cur.I][cur.J]
				queue = append(queue, newCur)
			}
		}
		visited[key] = struct{}{}
	}

	fmt.Println(dp[solution.N-1][solution.M-1])
}
