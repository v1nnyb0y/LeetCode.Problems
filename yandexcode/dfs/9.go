package dfs

import (
	"fmt"
	"log"
)

type Rewrite struct {
	N, M  int
	graph [][]int
}

func (r *Rewrite) inputFunc() {
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
	r.N, r.M = inputNumber(), inputNumber()
	r.graph = make([][]int, r.N+1)
	for i := 0; i < r.M; i++ {
		pair := inputArr(2)
		i, j := pair[0], pair[1]
		r.graph[i] = append(r.graph[i], j)
		r.graph[j] = append(r.graph[j], i)
	}
}

func (r *Rewrite) process() {
	visited := make([]int, r.N+1)
	hasCycle := false
	var dfs func(int, int)
	dfs = func(v, color int) {
		if hasCycle {
			return
		}
		visited[v] = color
		for _, vn := range r.graph[v] {
			if visited[vn] == 0 {
				dfs(vn, 3-color)
			} else {
				if visited[v] == visited[vn] {
					hasCycle = true
				}
			}
		}
	}

	for i := 1; i <= r.N; i++ {
		if hasCycle {
			break
		}
		if visited[i] == 0 {
			dfs(i, 1)
		}
	}

	if hasCycle {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func RewriteProcessing() {
	solution := &Rewrite{}
	solution.inputFunc()
	solution.process()
}
