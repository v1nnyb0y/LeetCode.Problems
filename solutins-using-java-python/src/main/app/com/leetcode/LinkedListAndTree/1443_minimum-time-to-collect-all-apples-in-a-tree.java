package com.leetcode.LinkedListAndTree;

import java.util.ArrayList;
import java.util.List;

class Solution_1443 {
    int ans;
    public int minTime(int n, int[][] edges, List<Boolean> hasApple) {
        ans = 0;
        List<Integer> []a = new ArrayList[n];
        for(int i=0;i<n;i++){
            a[i] = new ArrayList();
        }
        for(int []arr : edges){
            a[arr[0]].add(arr[1]);
            a[arr[1]].add(arr[0]);
        }
        boolean []visited = new boolean[n];
        visited[0] = true;
        rec(n, a, hasApple, visited, 0);
        return ans;
    }

    private boolean rec(int n, List<Integer> []a, List<Boolean> hasApple, boolean []visited, int ind){
        boolean flag = false;
        for(int i=0; i < a[ind].size(); i++){
            if(!visited[a[ind].get(i)]){
                visited[a[ind].get(i)] = true;
                boolean val = rec(n, a, hasApple, visited, a[ind].get(i));
                if(val){
                    ans += 2;
                }
                flag = flag | val;
            }
        }
        flag = flag | hasApple.get(ind);
        return flag;
    }
}
