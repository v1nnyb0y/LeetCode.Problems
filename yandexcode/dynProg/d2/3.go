package d2

import (
	"fmt"
	"log"
)

type Path struct {
	weight int
	path   string
}

type PathWithMaxWeight struct {
	N, M  int
	sizes []int
	graph [][]int
}

func (obj *PathWithMaxWeight) inputFunc() {
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
	obj.sizes = inputArr(2)
	obj.N, obj.M = obj.sizes[0], obj.sizes[1]
	obj.graph = make([][]int, obj.N)
	for i := 0; i < obj.N; i++ {
		obj.graph[i] = inputArr(obj.M)
	}
}

func (obj *PathWithMaxWeight) process() {
	dynamic := make([][]Path, obj.N+1)
	for i := range dynamic {
		dynamic[i] = make([]Path, obj.M+1)
	}

	for i := 0; i < obj.N; i++ {
		dynamic[i][0] = Path{weight: 0, path: ""}
	}
	for i := 0; i < obj.M; i++ {
		dynamic[0][i] = Path{weight: 0, path: ""}
	}

	for i := 1; i <= obj.N; i++ {
		var path = dynamic[i-1][1].path + " D"
		if i == 1 {
			path = ""
		}
		newPath := Path{
			weight: obj.graph[i-1][0] + dynamic[i-1][1].weight,
			path:   path,
		}
		dynamic[i][1] = newPath
	}
	for i := 1; i <= obj.M; i++ {
		var path = dynamic[1][i-1].path + " R"
		if i == 1 {
			path = ""
		}
		newPath := Path{
			weight: obj.graph[0][i-1] + dynamic[1][i-1].weight,
			path:   path,
		}
		dynamic[1][i] = newPath
	}

	for i := 2; i <= obj.N; i++ {
		for j := 2; j <= obj.M; j++ {
			val := obj.graph[i-1][j-1]
			var newPath Path
			if dynamic[i-1][j].weight+val <= dynamic[i][j-1].weight+val {
				newPath = Path{
					weight: dynamic[i][j-1].weight + val,
					path:   dynamic[i][j-1].path + " R",
				}
			} else {
				newPath = Path{
					weight: dynamic[i-1][j].weight + val,
					path:   dynamic[i-1][j].path + " D",
				}
			}
			dynamic[i][j] = newPath
		}
	}
	fmt.Println(dynamic[obj.N][obj.M].weight)
	fmt.Println(dynamic[obj.N][obj.M].path)
}

func PathWithMaxWeightProcess() {
	solution := PathWithMaxWeight{}
	solution.inputFunc()
	solution.process()
}
