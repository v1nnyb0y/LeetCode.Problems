package com.leetcode.Array;

import java.util.Arrays;
import java.util.TreeMap;

class Solution_2389 {
    public int[] answerQueries(int[] nums, int[] queries) {
        Arrays.sort(nums);
        TreeMap<Integer, Integer> map = new TreeMap<>();
        int cnt = 0;
        for(int i=0;i<nums.length;i++){
            map.put(cnt += nums[i], i+1);
        }
        int[] res = new int[queries.length];
        for(int i=0;i<queries.length;i++){
            res[i] = map.floorEntry(queries[i]) == null ? 0 : map.floorEntry(queries[i]).getValue();
        }
        return res;
    }
}
