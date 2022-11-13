import java.util.Arrays;

class Solution {
    public int pivotIndex(int[] nums) {
        int sumLeft = 0, sumRight = Arrays.stream(nums).sum();
        int result = -1;
        for(var i = 0; i < nums.length; i++) {
            if (sumLeft == sumRight - nums[i]) {
                result = i;
                break;
            }

            sumLeft += nums[i];
            sumRight -= nums[i];
        }
        return result;
    }
}