class NumMatrix {
    int[][] sum;
    int m,n;
    public NumMatrix(int[][] matrix) {
        int m = matrix.length+1;
        int n = matrix[0].length+1;
        sum = new int[m][n];
        for(int i=1;i<m;i++){
            for(int j=1;j<n;j++) sum[i][j]=sum[i-1][j] + sum[i][j-1] + matrix[i-1][j-1] - sum[i-1][j-1];
        }
    }

    public int sumRegion(int row1, int col1, int row2, int col2) {
        row1++;col1++;row2++;col2++;
        return (sum[row2][col2] - sum[row2][col1-1] - sum[row1-1][col2] + sum[row1-1][col1-1]);
    }
}