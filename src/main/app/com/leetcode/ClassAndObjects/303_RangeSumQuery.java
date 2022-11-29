package com.leetcode.ClassAndObjects;

class NumArray {
    private int[] current_nums;

    public NumArray(int[] nums) {
        current_nums = nums;
    }

    public int sumRange(int left, int right) {
        int sum = 0;
        for(var i = left; i <= right; i++) {
            sum += current_nums[i];
        }

        return sum;
    }
}