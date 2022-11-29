package com.leetcode.GraphTheory;

import javafx.util.Pair;

import java.util.*;

class Solution171 {
    public int[] shortestAlternatingPaths(int n, int[][] red_edges, int[][] blue_edges) {

        Map<Integer, List<int[]>> graph = new HashMap<>();
        for(int[] re: red_edges){
            graph.putIfAbsent(re[0], new ArrayList<>());
            graph.get(re[0]).add(new int[]{re[1],1});
        }
        for(int[] be: blue_edges){
            graph.putIfAbsent(be[0], new ArrayList<>());
            graph.get(be[0]).add(new int[]{be[1],-1});
        }
        Queue<int[]> q = new LinkedList<>();
        q.offer(new int[]{0,1});
        q.offer(new int[]{0,-1});
        int[] res = new int[n];
        Arrays.fill(res,-1);
        int dist=0;
        Set<Pair<Integer,Integer>> visited = new HashSet<>();
        visited.add(new Pair(0,1));
        visited.add(new Pair(0,-1));
        while(!q.isEmpty()){
            int l = q.size();
            for(int i=1;i<=l;i++){
                int[] x = q.poll();
                if(res[x[0]]==-1)
                    res[x[0]]=dist;
                for(int[] nei : graph.getOrDefault(x[0],new ArrayList<>())){
                    if(nei[1]==-1*x[1] && !visited.contains(new Pair(nei[0],nei[1]))){
                        q.offer(nei);
                        visited.add(new Pair(nei[0],nei[1]));
                    }
                }
            }
            dist++;
        }
        return res;
    }
}