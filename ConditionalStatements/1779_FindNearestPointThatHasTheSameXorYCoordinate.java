class Solution {
    public int nearestValidPoint(int x, int y, int[][] points) {
        int result = -1, minSum = Integer.MAX_VALUE;
        for (int i = 0; i < points.length; i++) {
            if (points[i][0] == x) {
                int manh = Math.abs(y - points[i][1]);
                if (minSum > manh) {
                    minSum = manh;
                    result = i;
                }
            } else if (points[i][1] == y) {
                int manh = Math.abs(x - points[i][0]);
                if (minSum > manh) {
                    minSum = manh;
                    result = i;
                }
            }
        }
        return result;
    }
}