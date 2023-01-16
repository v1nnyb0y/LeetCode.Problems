package com.leetcode.Array;

import java.util.ArrayList;
import java.util.List;

class Solution_57 {
    public int[][] insert(int[][] intervals, int[] newInterval) {

        List<int[]> result = new ArrayList<>();
        for (int[] i : intervals) {
            if (newInterval == null || i[1] < newInterval[0])
                result.add(i);
            else if (i[0] > newInterval[1]) {
                result.add(newInterval);
                result.add(i);
                newInterval = null;
            } else {
                newInterval[0] = Math.min(newInterval[0], i[0]);
                newInterval[1] = Math.max(newInterval[1], i[1]);
            }
        }
        if (newInterval != null)
            result.add(newInterval);
        return result.toArray(new int [0][]);
    }
}