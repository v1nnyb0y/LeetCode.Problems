package com.leetcode.Function;

class Solution91 {
    public boolean checkStraightLine(int[][] coordinates) {
        int size = coordinates.length;
        if (size == 2) {
            return true;
        }
        int x1 = coordinates[0][0], y1 = coordinates[0][1], x2 = coordinates[1][0], y2 = coordinates[1][1];
        for (int j = 2; j < size; j++) {
            int x3 = coordinates[j][0], y3 = coordinates[j][1];
            if ((y1 - y2) * (x1 - x3) != (y1 - y3) * (x1 - x2)) {
                return false;
            }
        }

        return true;
    }
}