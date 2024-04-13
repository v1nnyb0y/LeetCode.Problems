package com.leetcode.Strings

import java.util.Arrays

class `49_GroupAnagrams` {
    fun groupAnagrams(strs: Array<String>): List<List<String>> {
        val map = mutableMapOf<String, MutableList<String>>()
        strs.forEach {
            val key = it.toCharArray().sorted().joinToString("")
            if (!map.containsKey(key)) map[key] = mutableListOf()
            map[key]?.add(it)
        }

        return map.values.toList()
    }
}