package string

func WonderfulSubstrings(word string) int64 {
	get := func(c rune) int {
		return int(c - 'a')
	}

	cnt := make([]int64, 1024)
	cnt[0] = 1

	curState := 0
	res := int64(0)

	for _, c := range word {
		curState ^= 1 << get(c)

		res += cnt[curState]
		for odd := 'a'; odd <= 'j'; odd++ {
			oddState := curState ^ (1 << get(odd))
			res += cnt[oddState]
		}

		cnt[curState]++
	}

	return res
}
