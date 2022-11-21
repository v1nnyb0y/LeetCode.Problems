class Solution {
    public int maxScoreSightseeingPair(int[] arr) {
        int high_score = arr[0];
        int idx_i = 0;

        for (int idx_j = 1; idx_j < arr.length; ++idx_j) {
            int cur_score = (arr[idx_i] + arr[idx_j]) - (idx_j - idx_i);

            high_score = Math.max(
                    high_score,
                    cur_score
            );

            if (arr[idx_i] + idx_i < arr[idx_j] + idx_j) {
                idx_i = idx_j;
            }
        }

        return high_score;
    }
}