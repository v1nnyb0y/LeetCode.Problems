package com.leetcode.LinkedListAndTree;

import java.util.ArrayList;
import java.util.List;

class Solution_1519 {
    private int[] ans;
    private List<Integer>[] adj;
    public int[] countSubTrees(int n, int[][] edges, String labels) {
        adj = new List[n];
        for (int i = 0; i < n; i++) {
            adj[i] = new ArrayList<>();
        }
        for (int[] e : edges) {
            adj[e[0]].add(e[1]);
            adj[e[1]].add(e[0]);
        }

        ans = new int[n];
        dfs(0, -1, labels);
        return ans;
    }

    private int[] dfs(int currNode, int parent, String labels) {
        int[] currSubtreeFreq = new int[26];
        for (int child : adj[currNode]) {
            if (child == parent)  continue;
            int[] childSubtreeFreq = dfs(child, currNode, labels);
            for (int i = 0; i < 26; i++) {
                currSubtreeFreq[i] += childSubtreeFreq[i];
            }
        }
        currSubtreeFreq[labels.charAt(currNode) - 'a']++;
        ans[currNode] = currSubtreeFreq[labels.charAt(currNode) - 'a'];
        return currSubtreeFreq;
    }
}