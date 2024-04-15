package sort

import (
	"fmt"
	"log"
	"sort"
)

var input []int

func inputFunc() {
	input = make([]int, 3)
	for i := 0; i < len(input); i++ {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func MiddleElement() {
	inputFunc()
	sort.Ints(input)
	fmt.Println(input[1])
}
