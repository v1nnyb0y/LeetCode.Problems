package com.leetcode.Array;

class Solution36 {
    public int removeDuplicates(int[] nums) {
        int n = nums.length;
        if(n<=1){return n;}

        int left = 0;
        int right = 1;

        while(right<=n-1){
            if(nums[right] == nums[left]){
                right++;
            }
            else{
                nums[left+1] = nums[right];
                left++;
                right++;
            }
        }

        return left+1;
    }
}
