package bfs

import (
	"fmt"
	"log"
)

type PathInGraph struct {
	N, source, target int
	graph             [][]int
}

func (pig *PathInGraph) inputFunc() {
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

	pig.N = inputNumber()
	pig.graph = make([][]int, pig.N)
	for i := 0; i < pig.N; i++ {
		row := inputArr(pig.N)
		for j, cell := range row {
			if cell == 1 {
				pig.graph[i] = append(pig.graph[i], j)
			}
		}
	}
	pig.source, pig.target = inputNumber()-1, inputNumber()-1
}

func (pig *PathInGraph) process() {
	visited := make(map[int]int)

	visited[pig.target] = -1
	queue := []int{pig.target}

	var pathLen int
	for len(queue) > 0 {
		currLevelLen := len(queue)
		for i := 0; i < currLevelLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == pig.source {
				fmt.Println(pathLen)
				if pathLen != 0 {
					for node != -1 {
						fmt.Print(node+1, " ")
						prev := visited[node]
						node = prev
					}
				}
				return
			}

			for _, next := range pig.graph[node] {
				if _, ok := visited[next]; !ok {
					visited[next] = node
					queue = append(queue, next)
				}
			}
		}
		pathLen++
	}
	fmt.Println(-1)
}

func PathInGraphProcess() {
	solution := &PathInGraph{}
	solution.inputFunc()
	solution.process()
}
