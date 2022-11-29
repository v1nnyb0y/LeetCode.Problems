package com.leetcode.DynamicProgramming;

import java.util.List;

class Solution192 {
    public boolean wordBreak(String s, List<String> wordDict) {
        boolean dp[]=new boolean[s.length()+1];
        dp[0]=true;
        for(int i=0;i<=s.length();i++){
            int j=i-1;
            while(j>=0){
                if(dp[j]&&wordDict.contains(s.substring(j,i))){
                    dp[i]=true;
                    break;
                }
                j--;
            }
        }
        return dp[s.length()];
    }
}