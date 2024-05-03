package array

func GroupAnagrams(strs []string) [][]string {
	dict := make(map[[26]int][]string)
	for _, str := range strs {
		var key [26]int
		for _, chr := range str {
			key[chr-'a']++
		}
		dict[key] = append(dict[key], str)
	}
	var result [][]string
	for _, val := range dict {
		result = append(result, val)
	}
	return result
}
