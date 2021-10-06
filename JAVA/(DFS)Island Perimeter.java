//https://leetcode.com/problems/island-perimeter/

class Solution {
    int ans;
    int row;
    int col;
    public int islandPerimeter(int[][] grid) {
        row= grid.length;
        col = grid[0].length;
        ans = 0;
        
        for(int i = 0 ; i<row ; i++){
            boolean find = false;
            for(int j = 0 ; j<col ; j++){
                if(grid[i][j]==1){
                    dfs(grid, i ,j);
                    find = true;
                    break;
                }
            }
            if(find)    break;
        }
        return ans;
    }
    private void dfs(int[][] arr, int x, int y){
        arr[x][y] = -1;
        if(x-1<0)   ans++;
        else if(arr[x-1][y]==0) ans++;
        else if(arr[x-1][y]==1) dfs(arr, x-1, y);
        if(x+1>=row)    ans++;
        else if(arr[x+1][y]==0) ans++;
        else if(arr[x+1][y]==1) dfs(arr, x+1, y);
        
        if(y-1<0)   ans++;
        else if(arr[x][y-1]==0) ans++;
        else if(arr[x][y-1]==1) dfs(arr, x, y-1);
        if(y+1>=col)    ans++;
        else if(arr[x][y+1]==0) ans++;
        else if(arr[x][y+1]==1) dfs(arr, x, y+1);
    }
}