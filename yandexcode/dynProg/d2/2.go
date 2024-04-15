package d2

import (
	"fmt"
	"log"
)

var n, m int
var sizes []int
var graph [][]int

func inputFunc() {
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
	sizes = inputArr(2)
	n, m = sizes[0], sizes[1]
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = inputArr(m)
	}
}

func CheapestPath() {
	inputFunc()

	dynamic := make([][]int, n+1)
	for i := range dynamic {
		dynamic[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		dynamic[i][1] = graph[i-1][0] + dynamic[i-1][1]
	}
	for i := 1; i <= m; i++ {
		dynamic[1][i] = graph[0][i-1] + dynamic[1][i-1]
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			val := graph[i-1][j-1]
			dynamic[i][j] = min(dynamic[i-1][j]+val, dynamic[i][j-1]+val)
		}
	}
	fmt.Println(dynamic[n][m])
}
