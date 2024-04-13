package com.leetcode.Array;

import java.util.HashMap;
import java.util.HashSet;

class Solution1207 {
    boolean uniqueOccurrences(int[] arr) {
        var hashMap = new HashMap<Integer, Integer>();
        for(int val: arr) {
            hashMap.put(val, hashMap.getOrDefault(val, 0) + 1);
        }

        var hashSet = new HashSet<Integer>(hashMap.values());
        return hashSet.size() == hashMap.size();
    }
}