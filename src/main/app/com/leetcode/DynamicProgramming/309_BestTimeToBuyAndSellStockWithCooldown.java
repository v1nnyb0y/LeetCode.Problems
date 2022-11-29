package com.leetcode.DynamicProgramming;

class Solution203 {
    public int maxProfit(int[] prices) {
        int len = prices.length;
        if (len < 2) return 0;
        int[] hold   = new int[len];
        int[] unHold = new int[len];

        hold[0] = -prices[0];
        hold[1] = -Math.min(prices[0], prices[1]);
        unHold[1] = Math.max(0, hold[0] + prices[1]);

        for (int i = 2; i < len; i++){
            hold[i]   = Math.max(hold[i - 1], unHold[i - 2] - prices[i]);
            unHold[i] = Math.max(unHold[i - 1], hold[i - 1] + prices[i]);
        }

        return unHold[len - 1];
    }
}