package bfs

import (
	"fmt"
	"log"
)

type Blohi struct {
	N, M, S, T, Q int
	coordinates   map[[2]int]struct{}
}

func (b *Blohi) inputFunc() {
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
	b.N, b.M, b.S, b.T, b.Q = inputNumber(), inputNumber(), inputNumber(), inputNumber(), inputNumber()
	b.coordinates = make(map[[2]int]struct{})
	for i := 0; i < b.Q; i++ {
		pair := inputArr(2)
		b.coordinates[[2]int{pair[0] - 1, pair[1] - 1}] = struct{}{}
	}
}

func (b *Blohi) processing() {
	moves := [][]int{
		{-2, 1},
		{-2, -1},
		{1, 2},
		{-1, 2},
		{2, 1},
		{2, -1},
		{-1, -2},
		{1, -2},
	}
	var countOfBools int

	dp := make(map[[2]int]int)
	bfs := func(i, j int) {
		result := 0
		queue := [][2]int{{i, j}}

		for len(queue) > 0 {

			currLen := len(queue)
			for i := 0; i < currLen; i++ {
				coord := queue[0]
				queue = queue[1:]

				if _, ok := dp[coord]; ok {
					continue
				}
				dp[coord] = result

				if _, ok := b.coordinates[coord]; ok {
					countOfBools++
				}

				for _, move := range moves {
					newI := move[0] + coord[0]
					newJ := move[1] + coord[1]
					if newI >= b.N || newI < 0 || newJ >= b.M || newJ < 0 {
						continue
					}
					queue = append(queue, [2]int{newI, newJ})
				}
			}

			result++
			if countOfBools == b.Q {
				break
			}
		}
	}
	bfs(b.S-1, b.T-1)
	if countOfBools != b.Q {
		fmt.Println(-1)
	} else {
		sum := 0
		for coord := range b.coordinates {
			sum += dp[coord]
		}
		fmt.Println(sum)
	}
}

func BlohiProcessing() {
	solution := &Blohi{}
	solution.inputFunc()
	solution.processing()
}
