package dfs

import (
	"fmt"
	"log"
)

type FindCycle struct {
	N     int
	graph [][]int
}

func (fc *FindCycle) inputFunc() {
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

	fc.N = inputNumber()
	fc.graph = make([][]int, fc.N+1)
	for i := 1; i <= fc.N; i++ {
		row := inputArr(fc.N)
		for j := 0; j < len(row); j++ {
			if row[j] == 1 {
				fc.graph[j+1] = append(fc.graph[j+1], i)
			}
		}
	}
}

func (fc *FindCycle) process() {
	used := make([]int, fc.N+1)
	flag := 0
	var result []int

	var dfs func(int, int)
	dfs = func(v, p int) {
		if flag == 1 {
			return
		}
		used[v] = 1
		result = append(result, v)
		for _, to := range fc.graph[v] {
			if used[to] == 1 && to != p {
				result = append(result, to)
				flag = 1
				return
			} else if to != p {
				dfs(to, v)
			}
			if flag == 1 {
				return
			}
		}
		used[v] = 2
		result = result[:len(result)-1]
	}

	for i := 1; i <= fc.N; i++ {
		if used[i] == 0 {
			dfs(i, -1)
			if flag == 1 {
				break
			}
		}
	}

	// fmt.Println(result)
	if flag == 1 {
		i := len(result) - 2
		to := result[len(result)-1]
		for result[i] != to {
			i--
		}
		fmt.Println("YES")
		fmt.Println(len(result) - i - 1)
		for ; i < len(result)-1; i++ {
			fmt.Print(result[i], " ")
		}
	} else {
		fmt.Println("NO")
	}
}

func FindCycleProcess() {
	solution := &FindCycle{}
	solution.inputFunc()
	solution.process()
}
