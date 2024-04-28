package dfs

import (
	"fmt"
	"log"
)

type ComponentConnectivity struct {
	N, M  int
	graph [][]int
}

func (cc *ComponentConnectivity) inputFunc() {
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
		for i := 0; i < len(arr); i++ {
			_, err := fmt.Scan(&arr[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		return arr
	}

	cc.N, cc.M = inputNumber(), inputNumber()
	cc.graph = make([][]int, cc.N+1)
	for i := 0; i < cc.M; i++ {
		arr := inputArr(2)
		i, j := arr[0], arr[1]
		cc.graph[i] = append(cc.graph[i], j)
		cc.graph[j] = append(cc.graph[j], i)
	}
}

func (cc *ComponentConnectivity) process() {
	var level []int

	connectivity := make([][]int, 0)
	visited := make([]interface{}, cc.N+1)

	var dfs func(int)
	dfs = func(node int) {
		visited[node] = struct{}{}
		level = append(level, node)
		for _, next := range cc.graph[node] {
			if visited[next] == nil {
				dfs(next)
			}
		}
	}

	for start := 1; start <= cc.N; start++ {
		if visited[start] == nil {
			level = make([]int, 0)
			dfs(start)
			if len(level) != 0 {
				connectivity = append(connectivity, level)
			}
		}
	}

	fmt.Println(len(connectivity))
	for _, con := range connectivity {
		fmt.Println(len(con))
		for _, node := range con {
			fmt.Print(node, " ")
		}
		fmt.Println()
	}
}

func ComponentConnectivitySolution() {
	solution := &ComponentConnectivity{}
	solution.inputFunc()
	solution.process()
}
