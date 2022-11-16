class Solution {
    public int[][] matrixReshape(int[][] mat, int r, int c) {
        List<Integer> tempArray = new ArrayList<Integer>();
        int tempArrayPointer = 0;
        int[][] result = new int[r][c];
        int row = 0;
        int column = 0;

        int original_r = mat.length;
        int original_c = mat[0].length;
        if (original_c * original_r != r * c) {
            return mat;
        }

        for (int i = 0; i < original_r; i++) {
            for (int j = 0; j < original_c; j++) {
                tempArray.add(mat[i][j]);
            }
        }
        while (row < r) {
            while (column < c) {
                result[row][column] = tempArray.get(tempArrayPointer);
                tempArrayPointer += 1;
                column += 1;
            }
            column = 0;
            row += 1;
        }
        return result;
    }
}