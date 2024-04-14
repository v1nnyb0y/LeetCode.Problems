package string

var romanToIntMap = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func RomanToInt(s string) int {
	var result, i int
	for i < len(s) {
		if i+1 < len(s) {
			chr := string(s[i : i+2])
			if val, ok := romanToIntMap[chr]; ok {
				result += val
				i += 2
				continue
			}
		}

		chr := string(s[i])
		result += romanToIntMap[chr]
		i++
	}
	return result
}
