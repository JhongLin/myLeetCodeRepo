class Solution {
    public int[][] matrixReshape(int[][] mat, int r, int c) {
        int[][] ans = new int[r][c];
        int rows = mat.length, cols = mat[0].length;
        if(r*c!=rows*cols)  return mat;
        int nowCol=0;
        int nowRow=0;
        for (int i = 0 ; i<rows; i++) {
            
            for (int j = 0; j<cols ; j++) {
                if(nowCol>=c) {
                    nowRow++;
                    nowCol %= c;
                }
                
                ans[nowRow][nowCol] = mat[i][j];
                
                nowCol++;
            }
        }
        return ans;
    }
}