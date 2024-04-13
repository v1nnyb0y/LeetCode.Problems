package com.leetcode.BasicDataTypes

class `1523_CountOddNumbersInAnIntervalRange` {
    fun countOdds(low: Int, high: Int): Int {
        val fromVal: Int = if (low % 2 == 0) low + 1 else low
        var count = 0
        for (i in fromVal..high step 2) {
            count++
        }

        return count
    }
}
