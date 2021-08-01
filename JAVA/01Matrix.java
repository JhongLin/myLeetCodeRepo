class Solution {
    int m,n;
    public int[][] updateMatrix(int[][] mat) {
        m = mat.length; n = mat[0].length;
        int[][] ans = new int[m][n];
        for(int i = 0 ; i<m ; i++)
            Arrays.fill(ans[i], -1);
        for(int i = 0 ; i<m ; i++){
            for(int j = 0 ; j<n ; j++){
                if(mat[i][j]==0)    ans[i][j] = 0;
            }
        }
        boolean Done = false;
        int status = 0;
        while(!Done){
            Done = true;
            for(int i = 0 ; i<m ; i++){
                for(int j = 0 ; j<n ; j++){
                    if(ans[i][j]==-1){
                        bfs(mat, ans, i, j, status);
                        Done = false;
                    }
                }
            }
            status++;
        }
        return ans;
    }
    private void bfs(int[][] mat, int[][] ans, int a, int b,int cur){
        if(a-1>=0&&ans[a-1][b]==cur){ //up
            ans[a][b] = cur+1;
            return;
        }
        if(a+1<m&&ans[a+1][b]==cur){ //down
            ans[a][b] = cur+1;
            return;
        }
        if(b-1>=0&&ans[a][b-1]==cur){ //left
            ans[a][b] = cur+1;
            return;
        }
        if(b+1<n&&ans[a][b+1]==cur){ //right
            ans[a][b] = cur+1;
            return;
        }
    }
}