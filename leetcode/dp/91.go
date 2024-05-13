package dp

func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	// dp[i] := the number of ways to decode s[i..n)
	dp := make([]int, n+1)
	dp[n] = 1 // ""
	if s[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] += dp[i+1]
		}
		if s[i] == '1' || s[i] == '2' && s[i+1] < '7' {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}
