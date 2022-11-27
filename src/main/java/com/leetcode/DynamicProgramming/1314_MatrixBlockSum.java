package com.leetcode.DynamicProgramming;

class Solution184 {
    public int[][] matrixBlockSum(int[][] mat, int K) {
        int n = mat.length;
        int m = mat[0].length;
        int [][] dp = new int[n][m];
        for(int i = 0; i < n; i++){
            for(int j = 0; j < m; j++){
                if(i == 0 && j == 0){
                    dp[i][j] = mat[i][j];
                }else if(i == 0){
                    dp[i][j] = mat[i][j] + dp[i][j - 1];
                }else if(j == 0){
                    dp[i][j] = mat[i][j] + dp[i - 1][j];
                }else{
                    dp[i][j] = -dp[i - 1][j - 1] + dp[i - 1][j] + dp[i][j - 1] + mat[i][j];
                }
            }
        }
        int [][] result = new int[n][m];
        for(int i = 0; i < n; i++){
            for(int j = 0; j < m; j++){
                int ni = Math.min(i + K, n - 1);
                int nj = Math.min(j + K, m - 1);
                result[i][j] = dp[ni][nj];
                int ni2 = Math.max(i - K, 0);
                int nj2 = Math.max(j - K, 0);
                if(ni2 != 0 && nj2 != 0){
                    result[i][j] += dp[ni2 - 1][nj2 - 1];
                }
                if(ni2 != 0){
                    result[i][j] -= dp[ni2 - 1][nj];
                }
                if(nj2 != 0){
                    result[i][j] -= dp[ni][nj2 - 1];
                }
            }
        }
        return result;
    }
}