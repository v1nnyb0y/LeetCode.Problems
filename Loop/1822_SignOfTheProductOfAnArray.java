class Solution {
    public int arraySign(int[] nums) {
        int sign = 0;
        if (nums[0] == 0) return 0;
        if (nums[0] > 0) sign = 1;
        if (nums[0] < 0) sign = -1;

        for (var i = 1; i < nums.length; i++) {
            if (nums[i] == 0) return 0;
            if (sign == 1 && nums[i] > 0) continue;
            if (sign == 1 && nums[i] < 0) {
                sign = -1;
                continue;
            }
            if (sign == -1 && nums[i] > 0) continue;
            if (sign == -1 && nums[i] < 0) {
                sign = 1;
            }
        }

        return sign;
    }
}