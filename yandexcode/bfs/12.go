package bfs

import (
	"fmt"
	"log"
)

type ShortestPath struct {
	N, target, source int
	graph             [][]int
}

func (sp *ShortestPath) inputFunc() {
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

	sp.N = inputNumber()
	sp.graph = make([][]int, sp.N)
	for i := 0; i < sp.N; i++ {
		row := inputArr(sp.N)
		for j, cell := range row {
			if cell == 1 {
				sp.graph[i] = append(sp.graph[i], j)
			}
		}
	}
	sp.source, sp.target = inputNumber()-1, inputNumber()-1
}

func (sp *ShortestPath) process() {
	visited := make(map[int]struct{})
	queue := []int{sp.source}

	var shortestPath int
	for len(queue) > 0 {
		currLevelLen := len(queue)

		for i := 0; i < currLevelLen; i++ {
			node := queue[0]
			queue = queue[1:]
			visited[node] = struct{}{}

			if node == sp.target {
				fmt.Println(shortestPath)
				return
			}

			for _, next := range sp.graph[node] {
				if _, ok := visited[next]; !ok {
					queue = append(queue, next)
				}
			}
		}
		shortestPath++
	}

	fmt.Println(-1)
}

func ShortestPathProcess() {
	solution := &ShortestPath{}
	solution.inputFunc()
	solution.process()
}
