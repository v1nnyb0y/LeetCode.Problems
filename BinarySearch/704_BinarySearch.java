class Solution {
    public int search(int[] nums, int target) {
        int result = -1, left = 0, right = nums.length - 1;
        while (left <= right) {
            var pivot = left + ((right - left) / 2)
            if (nums[pivot] == target) {
                result = pivot;
                break;
            } else if (nums[pivot] > target) {
                right -= 1;
            } else {
                left += 1;
            }
        }
        return result;
    }
}