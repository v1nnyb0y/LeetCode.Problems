package com.leetcode.Array;

class Solution46 {
    public int[] runningSum(int[] nums) {
        int prevSum = nums[0];
        for(var i = 1; i < nums.length; i++) {
            var tmp = prevSum;
            prevSum += nums[i];
            nums[i] += tmp;
        }
        return nums;
    }
}