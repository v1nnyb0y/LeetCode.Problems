package com.leetcode.GraphTheory;

import java.util.LinkedList;
import java.util.List;

class Solution144 {
    public List<Integer> findSmallestSetOfVertices(int n, List<List<Integer>> edges) {
        int ver[] = new int[n];
        for(List<Integer> e : edges) ver[e.get(1)] = 1;

        List<Integer> list = new LinkedList<>();
        for(int i = 0; i != n; i++)
            if(ver[i] == 0) list.add(i);

        return list;
    }
}