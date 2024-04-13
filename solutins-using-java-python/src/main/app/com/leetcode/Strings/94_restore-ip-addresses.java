package com.leetcode.Strings;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

class Solution_94 {
    public List<String> restoreIpAddresses(String s) {
        List<String> result = new ArrayList<>();
        helper(s, 3, result, 0, new ArrayList<>());
        return result;

    }

    private void helper(String string, int remainDotToAdd, List<String> result, int startIndex, List<String> tempResult) {
        if (remainDotToAdd < 0) {
            if (startIndex >= string.length()) {
                result.add(tempResult.stream().map(String::toString).collect(Collectors.joining(".")));
                return;
            } else {
                return;
            }
        }

        for (int endIndex = startIndex + 1; endIndex <= string.length(); ++endIndex) {
            String sub = string.substring(startIndex, endIndex);
            int value = Integer.valueOf(sub);
            if (sub.equals(Integer.toString(value)) && validIpPart(value)) {
                tempResult.add(sub);
                helper(string, remainDotToAdd - 1, result, endIndex, tempResult);
                tempResult.remove(tempResult.size() - 1);
            } else {
                break;
            }
        }
    }

    private boolean validIpPart(int value) {
        return (value >= 0 && value <= 255);
    }
}