package array

func RemoveDuplicates(nums []int) int {
	k := 1
	potential := nums[0]
	for i := 1; i < len(nums); i++ {
		if potential != nums[i] {
			k++
			potential = nums[i]
		}
	}

	currElement := nums[0]
	i, j := 1, 1
	for i < k {
		for currElement == nums[j] {
			j++
		}
		nums[i] = nums[j]
		currElement = nums[j]
		i++
	}

	return k
}
