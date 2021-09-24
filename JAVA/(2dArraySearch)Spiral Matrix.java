class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> ans = new ArrayList<>();
        int i = 0, j = 0;
        char direction='r'; // 'r' 'b' 'l' 'a'
        while(i<matrix.length&&i>=0&&j>=0&&j<matrix[0].length&&matrix[i][j]!=101){
            ans.add(matrix[i][j]);
            matrix[i][j] = 101;
            if(direction=='r'){
                if(j+1<matrix[0].length&&matrix[i][j+1]!=101)  j++;
                else{
                    direction = 'b';
                    i++;
                }
            }else if(direction=='b'){
                if(i+1<matrix.length&&matrix[i+1][j]!=101)  i++;
                else{
                    direction = 'l';
                    j--;
                }
            }else if(direction=='l'){
                if(j-1>=0&&matrix[i][j-1]!=101)  j--;
                else{
                    direction = 'a';
                    i--;
                }
            }else {
                if(i-1>=0&&matrix[i-1][j]!=101)  i--;
                else{
                    direction = 'r';
                    j++;
                }
            }
        }
        return ans;
    }
}