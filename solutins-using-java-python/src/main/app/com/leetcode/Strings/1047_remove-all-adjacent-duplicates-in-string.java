package com.leetcode.Strings;

class Solution12 {
    public String removeDuplicates(String S) {
        int i = 0, j = 0;
        char[] res = S.toCharArray();
        for (; j < res.length; j++, i++) {
            res[i] = res[j];
            if (i > 0 && res[i-1] == res[j]){
                i -= 2;
            }
        }
        return new String(res, 0, i);
    }
}