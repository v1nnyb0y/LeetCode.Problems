package com.leetcode.DynamicProgramming

class `1143_LongestCommonSubsequence` {
    fun longestCommonSubsequence(text1: String, text2: String): Int {
        val dp = Array(text1.length + 1) { IntArray(text2.length + 1) }

        for (row in 1..text1.length) {
            val rowChar = text1[row - 1]

            for (col in 1..text2.length) {
                val colChar = text2[col - 1]

                if (rowChar == colChar) {
                    val previous = dp[row - 1][col - 1]
                    dp[row][col] = previous + 1
                } else {
                    val previousOnTop = dp[row - 1][col]
                    val previousOnLeft = dp[row][col - 1]
                    dp[row][col] = maxOf(previousOnTop, previousOnLeft)
                }
            }
        }

        return dp.last().last()
    }
}
