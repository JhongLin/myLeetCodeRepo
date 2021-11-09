//https://leetcode.com/problems/word-search/

class Solution {
    int m;
    int n;
    public boolean exist(char[][] board, String word) {
        m = board.length;
        n = board[0].length;
        char[] str = word.toCharArray();
        
        for(int i = 0 ; i<m ; i++ ){
            for(int j = 0 ; j<n ; j++ ){
                if(board[i][j]==str[0]){
                    if(str.length==1)   return true;
                    char temp = board[i][j];
                    board[i][j] = '0';
                    if(dfs(board, str, 1, i, j))    return true;
                    else    board[i][j] = temp;
                }
            }
        }
        
        return false;
    }
    private boolean dfs(char[][] board, char[] word, int index, int row, int col){
        if(row-1>=0&&board[row-1][col]==word[index]){
            if(index==word.length-1)    return true;
            char t = board[row-1][col];
            board[row-1][col] = '0';
            if(dfs(board, word, index+1, row-1, col))   return true;
            else    board[row-1][col] = t;
        }
        if(row+1<m&&board[row+1][col]==word[index]){
            if(index==word.length-1)    return true;
            char t = board[row+1][col];
            board[row+1][col] = '0';
            if(dfs(board, word, index+1, row+1, col))   return true;
            else    board[row+1][col] = t;
        }
        if(col-1>=0&&board[row][col-1]==word[index]){
            if(index==word.length-1)    return true;
            char t = board[row][col-1];
            board[row][col-1] = '0';
            if(dfs(board, word, index+1, row, col-1))   return true;
            else    board[row][col-1] = t;
        }
        if(col+1<n&&board[row][col+1]==word[index]){
            if(index==word.length-1)    return true;
            char t = board[row][col+1];
            board[row][col+1] = '0';
            if(dfs(board, word, index+1, row, col+1))   return true;
            else    board[row][col+1] = t;
        }
        return false;
    }
}