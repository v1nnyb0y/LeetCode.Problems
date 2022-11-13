class Solution {
    public int maxSubArray(int[] nums) {
        int n = nums.length;
        int[] maxSubArray = new int[n];
        maxSubArray[0] = nums[0];
        int max = maxSubArray[0];

        for(int i = 1; i < n; i++){
            maxSubArray[i] = nums[i] + (maxSubArray[i - 1] > 0 ? maxSubArray[i - 1] : 0);
            max = Math.max(max, maxSubArray[i]);
        }

        return max;
    }
}