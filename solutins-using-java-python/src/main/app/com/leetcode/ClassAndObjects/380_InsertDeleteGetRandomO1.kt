package com.leetcode.ClassAndObjects

class `380_InsertDeleteGetRandomO1`

class RandomizedSet {
    val hashSet = HashSet<Int>()

    fun insert(value: Int): Boolean {
        return hashSet.add(value)
    }

    fun remove(value: Int): Boolean {
        return hashSet.remove(value)
    }

    fun getRandom(): Int {
        return hashSet.random()
    }
}
