package com.leetcode.Array

class `1207_unique-number-of-occurrences` {
    fun uniqueOccurrences(arr: IntArray): Boolean {
        val hashMap: HashMap<Int, Int> = HashMap()
        for (value in arr) {
            hashMap[value] = hashMap[value] ?: 0
            hashMap[value] = hashMap[value]!! + 1
        }

        val hashSet: HashSet<Int> = HashSet()
        for (pair in hashMap) {
            hashSet.add(pair.value)
        }
        return hashSet.size == hashMap.size
    }
}