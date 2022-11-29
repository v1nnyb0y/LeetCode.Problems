package com.leetcode.BinarySearch;

class Solution59 {
    public int specialArray(int[] nums) {
        int X = -1;
        int low = 1, high = nums.length;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (countGreaterThanX(nums, mid) > mid) low = mid + 1;
            else if (countGreaterThanX(nums, mid) < mid) high = mid - 1;
            else {
                X = mid;
                break;
            }
        }
        return X;
    }

    public int countGreaterThanX(int[] nums, int X) {
        int res = 0;
        for (int num : nums) res += num >= X ? 1 : 0;
        return res;
    }
}