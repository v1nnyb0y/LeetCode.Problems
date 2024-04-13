package com.leetcode.BinarySearch;

class Solution54 {
    public int search(int[] nums, int target) {
        int n = nums.length;
        int lowIndex = 0, highIndex = n-1;

        while (lowIndex <= highIndex) {
            int midIndex =  (highIndex + lowIndex) / 2;

            if (nums[midIndex] == target) return midIndex;
            if (nums[lowIndex] <= nums[midIndex]) {
                if (target <= nums[midIndex] && target >= nums[lowIndex]) {
                    highIndex = midIndex - 1;
                } else {
                    lowIndex = midIndex + 1;
                }
            } else {
                if (target <= nums[highIndex] && target >= nums[midIndex]) {
                    lowIndex = midIndex + 1;
                } else {
                    highIndex = midIndex - 1;
                }
            }
        }

        return -1;
    }
}