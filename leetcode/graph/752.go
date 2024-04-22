package graph

func openLock(deadends []string, target string) int {
	q := make([]string, 0, 10000)
	q = append(q, "0000")
	min := 0

	dirs := map[rune][]rune{
		'0': {'1', '9'},
		'1': {'0', '2'},
		'2': {'1', '3'},
		'3': {'2', '4'},
		'4': {'3', '5'},
		'5': {'4', '6'},
		'6': {'5', '7'},
		'7': {'6', '8'},
		'8': {'7', '9'},
		'9': {'8', '0'},
	}

	visited := make(map[string]struct{}, 10000)
	visited["0000"] = struct{}{}
	for _, dead := range deadends {
		if dead == "0000" {
			return -1
		}
		visited[dead] = struct{}{}
	}

	for len(q) > 0 {
		level := []string{}
		for len(q) > 0 {
			lock := q[0]
			q = q[1:]

			if lock == target {
				return min
			}

			for i, dc := range lock {
				for _, next := range dirs[dc] {
					tryArr := []rune(lock)
					tryArr[i] = next
					try := string(tryArr)
					if _, ok := visited[try]; ok {
						continue
					}
					visited[try] = struct{}{}
					level = append(level, try)
				}
			}
		}
		if len(level) == 0 {
			return -1
		}
		q = append(q, level...)
		min++
	}

	return min
}
