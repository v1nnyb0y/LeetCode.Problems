package contest

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Crystall struct {
	Words []string
}

func (c *Crystall) inputFunc() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		c.Words = append(c.Words, scanner.Text())
	}
}

func (c *Crystall) outputFunc() {
	splittedWords := make([][]string, 0)
	for _, w := range c.Words {
		wordParts := make([]string, 0)

		var part = string(w[0])
		symbol := w[0]
		i := 1
		for i < len(w) {
			if symbol == w[i] {
				part += string(w[i])
			} else {
				wordParts = append(wordParts, part)
				symbol = w[i]
				part = string(w[i])
			}
			i++
		}

		wordParts = append(wordParts, part)
		splittedWords = append(splittedWords, wordParts)
	}

	var result string
	var error bool

	lenPart := len(splittedWords[0])
	for _, wordParts := range splittedWords {
		if lenPart != len(wordParts) {
			error = true
			break
		}
	}

	if !error {
		for i := 0; i < lenPart; i++ {
			partPosition := make([]string, 3)
			partPosition[0] = splittedWords[0][i]
			partPosition[1] = splittedWords[1][i]
			partPosition[2] = splittedWords[2][i]

			if partPosition[0][0] != partPosition[1][0] ||
				partPosition[1][0] != partPosition[2][0] ||
				partPosition[0][0] != partPosition[2][0] {
				error = true
				break
			}

			sort.Slice(
				partPosition, func(i, j int) bool {
					return len(partPosition[i]) < len(partPosition[j])
				},
			)

			result += partPosition[1]
		}
	}

	if error {
		fmt.Println("IMPOSSIBLE")
		return
	}
	fmt.Println(result)
}

func CrystallProcessing() {
	solution := &Crystall{}
	solution.inputFunc()
	solution.outputFunc()
}
