package com.leetcode.DynamicProgramming;

import java.util.Arrays;

class Solution298 {
    int[] dp1;
    int[] dp2;
    public int rob(int[] nums) {
        if(nums.length == 1){
            return nums[0];
        }

        dp1 = new int[nums.length];
        dp2 = new int[nums.length];

        Arrays.fill(dp1, -1);
        Arrays.fill(dp2, -1);

        int zero = robbing(nums, 0, dp1.length-2, dp1);
        int first = robbing(nums, 1, dp2.length-1, dp2);

        return Math.max(zero, first);
    }

    public int robbing(int[] nums, int index, int breaking, int[] dp){
        if(index > breaking) {
            return 0;
        }

        if(dp[index] == -1) {
            dp[index] = Math.max(robbing(nums, index+2, breaking, dp)+nums[index], robbing(nums, index+1, breaking, dp));
        }

        return dp[index];
    }
}