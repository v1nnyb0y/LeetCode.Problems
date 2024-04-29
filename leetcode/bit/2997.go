package bit

func MinOperations(nums []int, k int) int {
	bitwise := nums[0]
	for _, num := range nums[1:] {
		bitwise ^= num
	}

	counter := 0
	for bitwise != 0 || k != 0 {
		lastB, lastK := bitwise&1, k&1
		if lastB != lastK {
			counter++
		}

		bitwise >>= 1
		k >>= 1
	}
	return counter
}
