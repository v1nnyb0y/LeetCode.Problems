package com.leetcode.BasicDataTypes;

class Solution174 {
    public int countOdds(int low, int high) {
        int result = 0;
        if (low % 2 == 0) low += 1;
        for(var i = low; i<=high; i+=2) {
            result++;
        }
        return result;
    }
}