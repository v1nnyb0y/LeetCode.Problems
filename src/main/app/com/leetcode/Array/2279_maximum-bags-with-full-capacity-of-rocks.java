package com.leetcode.Array;

import java.util.Arrays;

class Solution_2279 {
    public int maximumBags(int[] capacity, int[] rocks, int additionalRocks) {
        int result = 0;
        int[] diff = new int[capacity.length];
        for(var i = 0; i < capacity.length; i++) {
            diff[i] = capacity[i] - rocks[i];
        }

        Arrays.sort(diff);
        for(var i = 0; i < diff.length; i++) {
            if (diff[i] == 0) {
                result += 1;
                continue;
            }

            if (additionalRocks <= 0) continue;
            additionalRocks -= diff[i];
            if (additionalRocks >= 0) {
                result += 1;
            }
        }

        return result;
    }
}