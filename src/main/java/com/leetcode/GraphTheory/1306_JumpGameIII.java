package com.leetcode.GraphTheory;

import java.util.HashSet;
import java.util.Set;

class Solution158 {
    Set<Integer> indexes = new HashSet<>();
    public boolean canReach(int[] arr, int start) {
        if (start < 0) return false;
        if (start > arr.length - 1) return false;
        if (indexes.contains(start)) return false;
        if (arr[start] == 0) return true;
        indexes.add(start);
        return canReach(arr, start + arr[start]) || canReach(arr, start - arr[start]);
    }
}