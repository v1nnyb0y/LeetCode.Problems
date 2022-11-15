class Solution {
    public int deleteAndEarn(int[] nums) {
        int a[]=new int[10001];
        for(int i:nums)
        {
            a[i]+=i;
        }
        int dp[]=new int[10001];
        dp[0]=a[0];
        dp[1]=Math.max(a[0],a[1]);
        for(int i=2;i<=10000;i++)
        {
            dp[i]=Math.max(dp[i-2]+a[i] , dp[i-1]);
        }
        return dp[10000];
    }
}