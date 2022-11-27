package com.leetcode.GraphTheory;

class Solution141 {
    public int numEnclaves(int[][] A) {
        int M = A.length, N = A[0].length;
        boolean[][] visited = new boolean[M][N];

        // starting from boundary
        for (int i = 0; i < M; ++i) {
            for (int j = 0; j < N; ++j) {
                if (i == 0 || i == M - 1 || j == 0 || j == N - 1) {
                    dfs(A, M, N, i, j, visited);
                }
            }
        }

        int count = 0;
        for (int i = 1; i < M - 1; ++i) {
            for (int j = 1; j < N - 1; ++j) {
                if (A[i][j] == 1 && !visited[i][j]) {
                    count++;
                }
            }
        }
        return count;
    }

    void dfs(int[][] A, int M, int N, int x, int y, boolean[][] visited) {
        if (x < 0 || y < 0 || x >= M || y >= N || visited[x][y] || A[x][y] == 0) {
            return;
        }

        visited[x][y] = true;
        dfs(A, M, N, x + 1, y, visited);
        dfs(A, M, N, x - 1, y, visited);
        dfs(A, M, N, x, y + 1, visited);
        dfs(A, M, N, x, y - 1, visited);
    }
}