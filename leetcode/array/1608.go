package array

import "fmt"

func specialArray(nums []int) int {
	for x := 0; x <= len(nums); x++ {
		counter := 0
		for _, num := range nums {
			if num >= x {
				counter++
			}
		}
		fmt.Println(x, counter)
		if counter == x {
			return x
		}
	}
	return -1
}
