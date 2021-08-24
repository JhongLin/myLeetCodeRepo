class Solution {
    public void setZeroes(int[][] matrix) {
        boolean col0 = false;
        int m = matrix.length, n = matrix[0].length;
        for(int i = 0 ; i<m ; i++){
            if(matrix[i][0]==0) col0=true;
            for(int j = 1 ; j <n ; j++){
                if(matrix[i][j]==0){
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }
        for(int i = m-1 ; i>=0 ; i--){
            for(int j = n-1 ; j>=1 ; j--){
                if(matrix[i][0]==0||matrix[0][j]==0){
                    matrix[i][j] = 0;
                }
            }
            if(col0)    matrix[i][0]=0;
        }
    }
}