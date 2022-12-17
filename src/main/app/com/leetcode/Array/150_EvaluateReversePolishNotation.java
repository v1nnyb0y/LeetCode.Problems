package com.leetcode.Array;

import java.util.*;

class Solution_150 {
    public int evalRPN(String[] tokens) {
        Stack<Integer> polish = new Stack<>();
        ArrayList<String> operations = new ArrayList<>(Arrays.asList("*", "+", "-", "/"));

        for (String token : tokens) {
            if (operations.contains(token)) {
                Integer second = polish.pop();
                Integer first = polish.pop();

                if (token.equals(operations.get(0))) {
                    polish.add(first * second);
                } else if (token.equals(operations.get(1))) {
                    polish.add(first + second);
                } else if (token.equals(operations.get(2))) {
                    polish.add(first - second);
                } else if (token.equals(operations.get(3))) {
                    polish.add((first - first % second) / second);
                }
            } else {
                polish.add(Integer.parseInt(token));
            }
        }

        return polish.peek();
    }
}