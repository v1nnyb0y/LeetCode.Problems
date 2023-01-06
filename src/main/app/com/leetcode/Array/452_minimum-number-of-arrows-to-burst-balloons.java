package com.leetcode.Array;

import java.util.Arrays;
import java.util.Comparator;

class Solution_452 {
    public int findMinArrowShots(int[][] points) {
        Arrays.sort(points, new Comparator<int[]>() {
            @Override
            public int compare(int[] x, int[] y) {
                if (x[0] > y[0]) return 1;
                else if (x[0] == y[0]) {
                    return x[1] - y[1];
                } else return -1;

            }
        });
        int count = 1;
        int prev = points[0][1];
        for (int i = 1; i < points.length; i++) {
            if (prev < points[i][0]) {
                prev = points[i][1];
                count++;
            } else prev = Math.min(prev, points[i][1]);
        }
        return count;
    }
}