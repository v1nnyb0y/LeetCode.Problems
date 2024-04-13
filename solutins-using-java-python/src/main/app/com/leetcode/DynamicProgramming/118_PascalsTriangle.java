package com.leetcode.DynamicProgramming;

import java.util.ArrayList;
import java.util.List;

class Solution204 {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> triangle = new ArrayList<List<Integer>>();
        triangle.add(new ArrayList<Integer>(1));
        triangle.get(0).add(1);

        for(int n = 2; n <= numRows; n++){
            ArrayList<Integer> row = new ArrayList<>(n);
            row.add(1);

            int i = 1;
            for(int lim = (n>>1) + (n%2); i != lim; i++)
                row.add(triangle.get(n-2).get(i) + triangle.get(n-2).get(i-1));

            for(int j = i - 1 - n%2; j >= 0; j--)
                row.add(row.get(j));

            triangle.add(row);
        }

        return triangle;
    }
}