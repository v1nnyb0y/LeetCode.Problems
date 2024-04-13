package com.leetcode.Array;

class Solution29 {
    public int diagonalSum(int[][] mat) {
        int sum = 0, len = mat.length;

        for (int i = 0; i < len; ++i) {
            sum += mat[i][i] + mat[i][len - i -1];
        }

        return len % 2 == 0 ? sum : sum - mat[len/2][len/2];
    }
}