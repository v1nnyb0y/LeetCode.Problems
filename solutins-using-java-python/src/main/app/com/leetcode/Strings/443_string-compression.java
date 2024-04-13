package com.leetcode.Strings;

class Solution_443 {
    public int compress(char[] chars) {
        int read = 0, write = 0;
        for (; read < chars.length;) {
            chars[write++] = chars[read++];
            int totalDups = 1;
            while (read < chars.length && chars[read] == chars[read-1]) {totalDups++; read++;}
            if (totalDups == 1) continue;

            int lenNeeded = Integer.toString(totalDups).length();
            for (int select = (int)Math.pow(10, lenNeeded-1); select >= 1; select /= 10) {
                chars[write++] = Character.forDigit(totalDups/select, 10);
                totalDups = totalDups % select;
            }
        }
        return write;
    }
}