package dfs

import (
	"fmt"
	"log"
)

type FindDepth struct {
	N, M  int
	graph [][]int
}

func (fd *FindDepth) inputFunc() {
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
	sizes := inputArr(2)
	fd.N, fd.M = sizes[0], sizes[1]
	fd.graph = make([][]int, fd.N)
	for i := 0; i < fd.N; i++ {
		fd.graph[i] = make([]int, 0, fd.M)
	}

	for i := 0; i < fd.M; i++ {
		pair := inputArr(2)
		l, r := pair[0], pair[1]
		fd.graph[l-1] = append(fd.graph[l-1], r)
		fd.graph[r-1] = append(fd.graph[r-1], l)
	}
}

func (fd *FindDepth) processing() {
	nodes := make([]interface{}, fd.N)
	var dfs func(int)
	dfs = func(node int) {
		nodes[node-1] = struct{}{}
		for _, next := range fd.graph[node-1] {
			if nodes[next-1] != nil {
				continue
			}
			dfs(next)
		}
	}

	dfs(1)

	counter := 0
	result := ""
	for node, ok := range nodes {
		if ok != nil {
			counter++
			result += fmt.Sprintf("%d ", node+1)
		}
	}
	fmt.Println(counter)
	fmt.Println(result)
}

func FindDepthProcessing() {
	solution := &FindDepth{}
	solution.inputFunc()
	solution.processing()
}
