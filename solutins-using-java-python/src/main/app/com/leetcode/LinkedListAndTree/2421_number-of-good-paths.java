package com.leetcode.LinkedListAndTree;

import java.util.Arrays;

class Solution_2421 {
    public int numberOfGoodPaths(int[] vals, int[][] edges) {
        Arrays.sort(edges, (a, b) -> {
            int tmp = Math.max(vals[a[0]], vals[a[1]]) - Math.max(vals[b[0]], vals[b[1]]);
            return tmp != 0 ? tmp : Math.min(vals[a[0]], vals[a[1]]) - Math.min(vals[a[0]], vals[a[1]]);
        });

        int res = vals.length;
        UnionFind3 uf = new UnionFind3(vals.length);
        for(int[] edge: edges)
            res += uf.union(edge[0], edge[1], vals);
        return res;
    }

    class UnionFind3 {
        private int[] data;
        private int[] count;

        public UnionFind3(int n){
            data = new int[n];
            count = new int[n];
            for (int i = 0; i <n; i++){
                data[i] = i;
                count[i] = 1;
            }
        }
        public int find(int x){
            if (data[x] == x) return x;
            return data[x] = find(data[x]);
        }

        public int union(int x, int y, int[] vals){
            int rootX = find(x);
            int rootY = find(y);
            if (rootX == rootY) return 0;

            int res = 0;
            if (vals[rootX] == vals[rootY]) {
                res = count[rootX] * count[rootY];

                count[rootX] += count[rootY];
                count[rootY] = 0;
            }

            if (vals[rootX] < vals[rootY])
                data[rootX] = rootY;
            else
                data[rootY] = rootX;

            return res;
        }
    }
}