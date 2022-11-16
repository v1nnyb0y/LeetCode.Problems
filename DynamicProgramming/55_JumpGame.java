class Solution {
    public boolean canJump(int[] nums) {
        if (nums == null || nums.length == 0) {
            return true;
        }
        if (nums.length == 1 && nums[0] == 0) {
            return true;
        }
        int len = nums.length;
        int far = 0;
        for (int i = 0; i < len; i++) {
            if (far < i) {
                break;
            }
            if (far < i + nums[i]) {
                far = i + nums[i];
            }
            if (far >= len - 1) {
                return true;
            }
        }
        return false;
    }
}