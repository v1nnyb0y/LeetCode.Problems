package com.leetcode.Array

class `503_continuous-subarray-sum` {
    fun checkSubarraySum(nums: IntArray, k: Int): Boolean {
        val hashMap: HashMap<Int, Int> = hashMapOf(0 to 0)
        var sum = 0
        for (i in nums.indices) {
            sum += nums[i]
            if (!hashMap.containsKey(sum % k)) {
                hashMap[sum % k] = i + 1
            } else if (hashMap[sum % k]!! < i) {
                return true
            }
        }

        return false
    }

}