package com.leetcode.BinarySearch;

class Solution50 {
    public int arrangeCoins(int n) {
        return (int)(Math.sqrt(2 * (long)n + 0.25) - 0.5);
    }
}