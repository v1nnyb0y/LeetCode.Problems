package emulate

import (
	"strconv"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	v1arr := strings.Split(version1, ".")
	v2arr := strings.Split(version2, ".")

	var i, j int
	for i < len(v1arr) || j < len(v2arr) {
		v1char, v2char := "0", "0"
		if i < len(v1arr) {
			v1char = v1arr[i]
			i++
		}

		if j < len(v2arr) {
			v2char = v2arr[j]
			j++
		}

		v1, _ := strconv.Atoi(v1char)
		v2, _ := strconv.Atoi(v2char)
		if v1 == v2 {
			continue
		}
		if v1 > v2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}
