package com.leetcode.BasicDataTypes

class `1491_AverageSalaryExcludingTheMinimumAndMaximumSalary` {
    fun average(salary: IntArray): Double {
        val max = salary.max()
        val min = salary.min()
        return (
            salary.sumOf {
                if (it == max || it == min) 0 else it
            } / (salary.size - 2)
            ).toDouble()
    }
}
