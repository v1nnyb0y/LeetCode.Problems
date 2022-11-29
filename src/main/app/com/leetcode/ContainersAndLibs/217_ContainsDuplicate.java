package com.leetcode.ContainersAndLibs;

import java.util.*;

class Solution97 {
    public boolean containsDuplicate(int[] nums) {
        Map<Integer, Integer> counter = new HashMap<Integer, Integer>();
        for(int num: nums) {
            if (counter.get(num) != null) {
                return true;
            } else {
                counter.put(num, 1);
            }
        }
        return false;
    }
}