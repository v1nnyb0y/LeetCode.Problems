package sort

import (
	"fmt"
	"log"
)

type TopSort struct {
	N, M  int
	graph [][]int
}

func (ts *TopSort) inputFunc() {
	inputArr := func(lenArr int) []int {
		var arr = make([]int, lenArr)
		for i := 0; i < len(arr); i++ {
			_, err := fmt.Scan(&arr[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		return arr
	}
	pair := inputArr(2)
	ts.N, ts.M = pair[0], pair[1]
	ts.graph = make([][]int, ts.N+1)
	for i := 0; i < ts.M; i++ {
		pair = inputArr(2)
		i, j := pair[0], pair[1]
		ts.graph[i] = append(ts.graph[i], j)
	}
}

func (ts *TopSort) processing() {
	colors := make([]int, ts.N+1)
	stack := make([]int, 0)
	error := false

	var dfs func(int)
	dfs = func(v int) {
		if error {
			return
		}
		colors[v] = 1
		for i := 0; i < len(ts.graph[v]); i++ {
			if colors[ts.graph[v][i]] == 0 {
				dfs(ts.graph[v][i])
			} else if colors[ts.graph[v][i]] == 1 {
				error = true
				return
			}
		}
		stack = append(stack, v)
		colors[v] = 2
	}

	for i := 1; i <= ts.N; i++ {
		if colors[i] == 0 {
			dfs(i)
		}
	}

	if error {
		fmt.Println(-1)
	} else {
		for i := len(stack) - 1; i >= 0; i-- {
			fmt.Print(stack[i], " ")
		}
	}
}

func TopSortProcessing() {
	solution := &TopSort{}
	solution.inputFunc()
	solution.processing()
}
