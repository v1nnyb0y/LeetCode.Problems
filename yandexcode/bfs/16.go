package bfs

import (
	"fmt"
	"log"
)

type Metro struct {
	N, M, source, target int
	metro, stationToLine [][]int
}

func (m *Metro) inputFunc() {
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

	m.N, m.M = inputNumber(), inputNumber()
	m.metro = make([][]int, m.M+1)
	m.stationToLine = make([][]int, m.N+1)
	for i := 1; i <= m.M; i++ {
		sc := inputNumber()
		stations := inputArr(sc)
		m.metro[i] = stations
		for _, s := range stations {
			if m.stationToLine[s] == nil {
				m.stationToLine[s] = []int{i}
			} else {
				m.stationToLine[s] = append(m.stationToLine[s], i)
			}
		}
	}
	m.source, m.target = inputNumber(), inputNumber()
}

func (m *Metro) outputFunc() {
	targetLines := m.stationToLine[m.target]
	findMap := make(map[int]struct{})
	for _, line := range targetLines {
		findMap[line] = struct{}{}
	}
	metroToCount := make([]int, m.M+1)
	visited := make([]bool, m.M+1)

	output := true
	queue := m.stationToLine[m.source]
	for _, line := range queue {
		visited[line] = true
	}

	for len(queue) > 0 {
		line := queue[0]
		queue = queue[1:]

		if _, ok := findMap[line]; ok {
			fmt.Println(metroToCount[line])
			output = false
			break
		}

		for _, station := range m.metro[line] {
			if len(m.stationToLine[station]) > 1 {
				for _, newLine := range m.stationToLine[station] {
					if !visited[newLine] {
						visited[newLine] = true
						queue = append(queue, newLine)
						if metroToCount[newLine] != 0 {
							metroToCount[newLine] = min(metroToCount[newLine], metroToCount[line]+1)
						} else {
							metroToCount[newLine] = metroToCount[line] + 1
						}
					}
				}
			}
		}
	}

	if output {
		fmt.Println(-1)
	}
}

func MetroProcessing() {
	solution := &Metro{}
	solution.inputFunc()
	solution.outputFunc()
}
