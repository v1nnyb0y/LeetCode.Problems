package com.leetcode.Loop;

import java.util.Arrays;

class Solution211 {
    public boolean canMakeArithmeticProgression(int[] arr) {
        arr = Arrays.stream(arr).sorted().toArray();
        int diff = arr[1] - arr[0];
        for(var i = 2; i < arr.length; i++) {
            if (diff != arr[i] - arr[i - 1]) return false;
        }

        return true;
    }
}