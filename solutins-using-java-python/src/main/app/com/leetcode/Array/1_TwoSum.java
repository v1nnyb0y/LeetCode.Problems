package com.leetcode.Array;

import java.util.HashMap;
import java.util.Map;

class Solution30 {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        for(var i = 0; i < nums.length; i++) {
            int num = nums[i];
            if (map.get(target - num) != null) {
                return new int[] {map.get(target - num), i};
            } else {
                map.put(num, i);
            }
        }
        return null;
    }
}