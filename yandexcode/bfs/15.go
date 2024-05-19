package bfs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type PathSteleolog struct {
	N       int
	current [3]int
	cave    [][][]int
}

func (ps *PathSteleolog) inputFunc() {
	inputNumber := func() int {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
	inputText := func() []string {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		var words []string
		for scanner.Scan() {
			text := scanner.Text()
			if text == "exit" {
				break
			}
			words = append(words, text)
		}
		return words
	}
	ps.N = inputNumber()

	ps.cave = make([][][]int, ps.N)
	for i := range ps.cave {
		ps.cave[i] = make([][]int, ps.N)
		for j := range ps.cave[i] {
			ps.cave[i][j] = make([]int, ps.N)
		}
	}

	words := inputText()

	counter := 0
	for level := 0; level < ps.N; level++ {
		for row := 0; row < ps.N; row++ {
			word := words[counter]
			counter++
			for ind, chr := range word {
				sChr := string(chr)
				if sChr == "S" {
					ps.current = [3]int{level, row, ind}
				}

				if sChr == "#" {
					ps.cave[level][row][ind] = 1
				} else {
					ps.cave[level][row][ind] = 0
				}
			}
		}
	}
}

func (ps *PathSteleolog) outputFunc() {
	moveX := [6]int{0, 0, 0, 0, -1, 1}
	moveY := [6]int{0, 0, -1, 1, 0, 0}
	moveZ := [6]int{-1, 1, 0, 0, 0, 0}
	steps := -1
	queue := [][3]int{ps.current}

	init := 2
	ps.cave[ps.current[0]][ps.current[1]][ps.current[2]] = init

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == 0 {
			steps = ps.cave[cur[0]][cur[1]][cur[2]]
			break
		}

		for i := 0; i < 6; i++ {
			z := cur[0] + moveZ[i]
			x := cur[1] + moveX[i]
			y := cur[2] + moveY[i]
			if z < 0 || x < 0 || y < 0 || z == ps.N || x == ps.N || y == ps.N {
				continue
			}
			if ps.cave[z][x][y] == 0 {
				ps.cave[z][x][y] = ps.cave[cur[0]][cur[1]][cur[2]] + 1
				queue = append(queue, [3]int{z, x, y})
			}
		}
	}
	fmt.Println(steps - init)
}

func PathSteleologProcessing() {
	solution := &PathSteleolog{}
	solution.inputFunc()
	solution.outputFunc()
}
