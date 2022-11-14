import java.util.Arrays;

class Solution {
    public int[] sortedSquares(int[] nums) {
        return Arrays.stream(nums).map(i->i*i).sorted().toArray();
    }
}