package com.leetcode.Array;

import java.util.Arrays;

class Solution_1833 {
    public int maxIceCream(int[] costs, int coins) {
        Arrays.sort(costs);

        int result = 0;
        for(int cost: costs) {
            if (cost > coins) break;
            coins -= cost;
            result += 1;
        }
        return result;
    }
}
