class Solution {
    public int minCostClimbingStairs(int[] cost) {
        if(cost.length==0)
            return 0;
        if(cost.length==1)
            return cost[0];
        int n=cost.length;
        int[]dp=new int[n+1];

        dp[0]=0;
        dp[1]=cost[0];
        for(int i=2;i<=n;i++)
        {
            dp[i]= cost[i-1] + Math.min(dp[i-1], dp[i-2]);
        }

        return Math.min(dp[n], dp[n-1]);
    }
}