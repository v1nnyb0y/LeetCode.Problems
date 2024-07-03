package contest

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type TwiceLetters struct {
	Words []string
}

func (tl *TwiceLetters) inputFunc() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		str := scanner.Text()
		//if str == "exit" {
		//	break
		//}

		if len(str) >= 2 {
			tl.Words = append(tl.Words, str)
		}
	}
}

func (tl *TwiceLetters) outputFunc() {
	counter := make(map[string]int)
	for _, word := range tl.Words {
		for i := 2; i <= len(word); i++ {
			counter[word[i-2:i]]++
		}
	}

	maxCounter := math.MinInt
	maxStr := ""

	for str, count := range counter {
		if count > maxCounter {
			maxStr = str
			maxCounter = count
		} else if count == maxCounter && maxStr < str {
			maxStr = str
		}
	}

	fmt.Println(maxStr)
}

func TwiceLettersProcessing() {
	solution := &TwiceLetters{}
	solution.inputFunc()
	solution.outputFunc()
}
