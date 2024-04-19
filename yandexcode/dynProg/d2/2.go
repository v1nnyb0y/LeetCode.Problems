package d2

import (
	"fmt"
	"log"
)

type CheapestPath struct {
	N, M  int
	sizes []int
	graph [][]int
}

func (cp *CheapestPath) inputFunc() {
	inputArr := func(size int) []int {
		var arr = make([]int, size)
		for i := 0; i < len(arr); i++ {
			_, err := fmt.Scan(&arr[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		return arr
	}
	cp.sizes = inputArr(2)
	cp.N, cp.M = cp.sizes[0], cp.sizes[1]
	cp.graph = make([][]int, cp.N)
	for i := 0; i < cp.N; i++ {
		cp.graph[i] = inputArr(cp.M)
	}
}

func CheapestPathProcess() {
	solution := &CheapestPath{}
	solution.inputFunc()

	dynamic := make([][]int, solution.N+1)
	for i := range dynamic {
		dynamic[i] = make([]int, solution.M+1)
	}

	for i := 1; i <= solution.N; i++ {
		dynamic[i][1] = solution.graph[i-1][0] + dynamic[i-1][1]
	}
	for i := 1; i <= solution.M; i++ {
		dynamic[1][i] = solution.graph[0][i-1] + dynamic[1][i-1]
	}

	for i := 2; i <= solution.N; i++ {
		for j := 2; j <= solution.M; j++ {
			val := solution.graph[i-1][j-1]
			dynamic[i][j] = min(dynamic[i-1][j]+val, dynamic[i][j-1]+val)
		}
	}
	fmt.Println(dynamic[solution.N][solution.M])
}
