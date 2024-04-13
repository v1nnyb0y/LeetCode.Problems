package com.leetcode.Array;

import java.util.Collections;
import java.util.PriorityQueue;

class Solution_1962 {
    public int minStoneSum(int[] piles, int k) {
        int totalStones = 0;
        PriorityQueue<Integer> pQueue = new PriorityQueue<>(Collections.reverseOrder());
        for(var pile: piles) {
            totalStones += pile;
            pQueue.add(pile);
        }

        while (k != 0) {
            var pile = pQueue.poll();

            int floor = (int)Math.floor(pile / 2);
            pile -= floor;
            totalStones -= floor;
            System.out.println(pile);
            pQueue.add(pile);
            k -= 1;
        }

        return totalStones;
    }
}