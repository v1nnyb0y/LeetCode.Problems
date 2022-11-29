package com.leetcode.Strings;

import java.util.ArrayList;

class Solution18 {
    public String reverseWords(String s) {
        String[] sArr = s.trim().split(" ");
        ArrayList<String> result = new ArrayList<String>();
        for (int i = sArr.length - 1, j = 0; i >= 0; i--, j++) {
            if (sArr[i] == "") continue;
            result.add(sArr[i]);
        }
        return String.join(" ", result);
    }
}