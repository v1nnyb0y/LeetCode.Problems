package com.leetcode.BinarySearch;

class Solution57 {
    public char nextGreatestLetter(char[] letters, char target) {

        for(Character ch : letters) {
            if(ch>target)
                return ch;
        }

        return letters[0];
    }
}