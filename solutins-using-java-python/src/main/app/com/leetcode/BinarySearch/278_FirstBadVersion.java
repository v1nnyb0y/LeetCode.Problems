package com.leetcode.BinarySearch;/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */

class Solution111 {
    public boolean isBadVersion(int version) {
        return true;
    }

    public int firstBadVersion(int n) {
        int left = 1, right = n;
        while (left != right) {
            int pivot = left + (right - left) / 2;
            if (isBadVersion(pivot)) {
                right = pivot;
            } else {
                left = pivot + 1;
            }
        }
        return right;
    }
}