package com.leetcode.Array;

import java.util.HashMap;
import java.util.Map;

class Solution_2244 {
    public int minimumRounds(int[] tasks)
    {
        Map<Integer, Integer> map = new HashMap<>();
        for(int i : tasks)
            map.put(i,map.getOrDefault(i,0)+1);
        int count = 0;
        for(int i : map.values())
        {
            if(i < 2)
                return -1;
            count += i/3;
            if(i%3 != 0)
                count++;
        }
        return count;
    }
}