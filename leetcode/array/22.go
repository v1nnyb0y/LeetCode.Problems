package array

func GenerateParenthesis(n int) []string {

	var stack string = ""
	var result []string

	var backtrack func(int, int)
	backtrack = func(start, close int) {
		if start == n && close == n {
			result = append(result, stack)
			return
		}

		if start < n {
			stack += "("
			backtrack(start+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < start {
			stack += ")"
			backtrack(start, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}
