package com.leetcode.GraphTheory;

import java.util.*;
import java.util.stream.Collectors;

class Solution165 {

    public int minimumJumps(int[] f, int a, int b, int x) {
        Set<Integer> forbidden = Arrays.stream(f).boxed().collect(Collectors.toSet());

        Queue<int[]> queue = new ArrayDeque<>();

        queue.offer(new int[]{0, 0, 1});    // currentPosition, cost, canJumpBack

        while (!queue.isEmpty()){
            int[] t = queue.poll();
            int pos = t[0];
            int cost = t[1];
            int canJumpBack = t[2];

            if(pos == x) return cost;
            if(forbidden.contains(pos)) continue;
            forbidden.add(pos);

            int next = pos + a;
            int prev = pos - b;

            if(!forbidden.contains(prev) && prev >= 0 && canJumpBack == 1) {
                queue.offer(new int[]{prev, cost + 1, 0});
            }

            if(!forbidden.contains(next) && next < 6000) {
                queue.offer(new int[]{next, cost + 1, 1});
            }
        }
        return -1;
    }
}