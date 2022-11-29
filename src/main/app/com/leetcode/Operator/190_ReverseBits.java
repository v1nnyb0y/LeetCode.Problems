package com.leetcode.Operator;

class Solution207 {
    // you need treat n as an unsigned value
    public int reverseBits(int n) {
        int res = 0, mask = 1;
        for(int i=0; i < 32; i++){
            res = res*2 + (n & mask);
            n = n >>> 1;
        }
        return res;
    }
}