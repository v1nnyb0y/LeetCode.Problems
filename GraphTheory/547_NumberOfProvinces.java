public class Solution {
    private void helper(int[][] matrix, boolean[][] visited, int i, int j){
        visited[i][j]=true;

        for(int col=0; col<matrix.length; col++){
            if( matrix[i][col]==1 && !visited[i][col] )
                helper(matrix, visited, i, col);

            if( matrix[j][col]==1 && !visited[j][col] )
                helper(matrix, visited, j, col);
        }
    }

    public int findCircleNum(int[][] M) {
        if(M.length==0)
            return 0;

        int n=M.length;
        int count=0;
        boolean[][] visited = new boolean[n][n];
        for(int i=0; i<n; i++)
            for(int j=0; j<n; j++)
                if(M[i][j]==1 && !visited[i][j]){
                    count++;
                    helper(M, visited, i, j);
                }

        return count;
    }
}
