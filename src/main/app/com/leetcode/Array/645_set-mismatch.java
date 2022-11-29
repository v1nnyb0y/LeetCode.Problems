package com.leetcode.Array;

class Solution42 {

    private static final int NOT_FOUND = -1;

    public int[] findErrorNums(int[] input) {
        int duplicate = NOT_FOUND;
        int missing = NOT_FOUND;

        for (int i = 0; i < input.length; ++i) {
            if (input[Math.abs(input[i]) - 1] < 0) {
                duplicate = Math.abs(input[i]);
                continue;
            }
            input[Math.abs(input[i]) - 1] = -input[Math.abs(input[i]) - 1];
        }

        for (int i = 0; i < input.length && missing == NOT_FOUND; ++i) {
            if (input[i] > 0) {
                missing = i + 1;
            }
        }
        return new int[]{duplicate, missing};
    }
}