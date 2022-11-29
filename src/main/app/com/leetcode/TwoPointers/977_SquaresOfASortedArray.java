package com.leetcode.TwoPointers;

import java.util.Arrays;

class Solution25 {
    public int[] sortedSquares(int[] nums) {
        return Arrays.stream(nums).map(i->i*i).sorted().toArray();
    }
}