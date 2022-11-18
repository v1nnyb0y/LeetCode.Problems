public class Solution {
    public int maxProduct(int[] nums) {
        if(nums.length<1) return 0;
        int max = nums[0];
        int maxPositive = 1;
        int maxNegative = 1;
        for(int i=0, L=nums.length; i<L; i++) {
            if(nums[i]>0) {
                maxPositive = maxPositive * nums[i];
                maxNegative = (maxNegative < 0) ? maxNegative * nums[i] : 1;
                max = maxPositive > max ? maxPositive : max;
            }
            else if(nums[i]<0) {
                int temp = maxNegative;
                maxNegative = maxPositive * nums[i];
                maxPositive = (temp < 0) ? temp * nums[i] : 1;
                max = (temp<0) ? (maxPositive > max ? maxPositive : max) : max;
            }
            else {
                maxPositive = maxNegative = 1;
                max = max > 0 ? max : 0;
            }
        }
        return max;
    }
}