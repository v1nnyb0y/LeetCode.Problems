package com.leetcode.Strings;

class Solution_125 {
    public boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        var arr = s.toCharArray();

        while (left < right) {
            char leftCh = arr[left];
            char rightCh = arr[right];

            if (Character.isLetterOrDigit(leftCh) && Character.isLetterOrDigit(rightCh)) {
                if (Character.toLowerCase(leftCh) != Character.toLowerCase(rightCh)) {
                    return false;
                }
                left += 1;
                right -= 1;
                continue;
            }

            if (!Character.isLetterOrDigit(leftCh)) {
                left += 1;
            }

            if (!Character.isLetterOrDigit(rightCh)) {
                right -= 1;
            }
        }
        return true;
    }
}