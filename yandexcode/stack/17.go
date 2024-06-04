package stack

import (
	"fmt"
	"log"
)

type TransportLine struct {
	N      int
	Answer []bool
}

func (tl *TransportLine) inputFunc() {
	inputNumber := func() int {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
	inputArr := func(lenArr int) []float64 {
		var arr = make([]float64, lenArr)
		for i := 0; i < len(arr); i++ {
			_, err := fmt.Scan(&arr[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		return arr
	}

	tl.N = inputNumber()
	for i := 0; i < tl.N; i++ {
		lenTest := inputNumber()
		test := inputArr(lenTest)
		var answer bool
		for j := 0; j < lenTest-2; j++ {
			if test[j+2] <= test[j] && test[j] < test[j+1] {
				answer = true
				break
			}
		}
		tl.Answer = append(tl.Answer, answer)
	}
}

func (tl *TransportLine) outputFunc() {
	for _, ans := range tl.Answer {
		if ans {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}

func TransportLineProcessing() {
	solution := &TransportLine{}
	solution.inputFunc()
	solution.outputFunc()
}
