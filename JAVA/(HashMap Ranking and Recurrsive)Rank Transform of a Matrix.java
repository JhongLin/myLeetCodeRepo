class Solution {
    int[][] ans;
    int[][] yokoRank;
    int[][] tatteRank;
    int m,n;
    public int[][] matrixRankTransform(int[][] matrix) {
        if(matrix==null)    return null;
        m = matrix.length; n = matrix[0].length;
        ans = new int[m][n];
        yokoRank = new int[m][n];
        tatteRank = new int[m][n];
        rowRanking(matrix);
        colRanking(matrix);
        for(int i=0; i<m; i++)  System.out.println(Arrays.toString(matrix[i]));
        //for(int i=0; i<m; i++)  System.out.println(Arrays.toString(yokoRank[i]));
        //System.out.println("");
        //for(int i=0; i<m; i++)  System.out.println(Arrays.toString(tatteRank[i]));
        //ans[0] = new int[n];
        //ans[1] = new int[n];
        /*
        for(int i = 0 ; i<m ; i++){
            Arrays.fill(ans[i], Integer.MIN_VALUE);
        }*/
        
        for(int i = 0 ; i<m ; i++){
            for(int j = 0 ; j<n ; j++){
                if(ans[i][j]==0){
                    if(yokoRank[i][j]==1&&yokoRank[i][j]==tatteRank[i][j]){
                        ans[i][j] = 1;
                        continue;
                    }
                    if(yokoRank[i][j]>tatteRank[i][j]){
                        ans[i][j] = yokoRank[i][j];
                        validateCol(i, j, yokoRank[i][j]);
                    }else if(yokoRank[i][j]<tatteRank[i][j]){
                        ans[i][j] = tatteRank[i][j];
                        validateRow(i, j, tatteRank[i][j]);
                    }else{
                        ans[i][j] = yokoRank[i][j];
                        //validateRow(i, j, yokoRank[i][j]);
                    }
                }
            }
        }
        /*
        for(int i = 0 ; i<m ; i++){
            for(int j = 0 ; j<n ; j++){
                for(int k = 0 ; k<m ; k++){
                    if(tatteRank[i][j]<tatteRank[k][j]&&ans[i][j]>=ans[k][j]){
                        ans[k][j] = ans[i][j]+tatteRank[k][j]-tatteRank[i][j];
                        validateCol(i, j, ans[i][j]+tatteRank[k][j]-tatteRank[i][j]);
                    }
                }
                for(int k = 0 ; k<n ; k++){
                    
                }
            }
        }
        */
        //System.out.println("");
        //for(int i=0; i<m; i++)  System.out.println(Arrays.toString(ans[i]));
        return ans;
    }
    private void rowRanking(int[][] matrix){
        for(int i = 0 ; i<m ; i++){
            int[] temp = Arrays.copyOfRange(matrix[i], 0, n);
            Arrays.sort(temp);
            HashMap<Integer, Integer> map = new HashMap<>();
            int rank = 1;
            for(int j = 0; j<n ; j++){
                if(map.get(temp[j])==null)  map.put(temp[j], rank++);
            }
            for(int j = 0;j<n;j++){
                yokoRank[i][j] = map.get(matrix[i][j]);
            }
        }
    }
    private void colRanking(int[][] matrix){
        for(int i = 0 ; i<n ; i++){
            int[] temp = new int[m];
            for(int j = 0 ; j<m ; j++)  temp[j] = matrix[j][i];
            Arrays.sort(temp);
            HashMap<Integer, Integer> map = new HashMap<>();
            int rank = 1;
            for(int j = 0; j<m ; j++){
                if(map.get(temp[j])==null)  map.put(temp[j], rank++);
            }
            for(int j = 0;j<m;j++){
                tatteRank[j][i] = map.get(matrix[j][i]);
            }
        }
    }
    
    private void validateRow(int x, int y, int after){
        //if(x==3&&y==8)  System.out.println("Call vRow at "+Integer.toString(x)+" "+Integer.toString(y)+" with "+Integer.toString(after));
        ans[x][y] = after;
        /*
        if(after<tatteRank[x][y]){
            ans[x][y] = tatteRank[x][y];
            validateRow(x, y, tatteRank[x][y]);
            //validateCol(x, y, tatteRank[x][y]);
            return;
        }*/
        
        if(after<yokoRank[x][y]){
            ans[x][y] = yokoRank[x][y];
            //validateRow(x, y, yokoRank[x][y]);
            validateCol(x, y, yokoRank[x][y]);
            return;
        }
        int orig = yokoRank[x][y];
        int diff = after - orig;
        for(int i = 0 ; i<n ; i++){
            if(i==y)    continue;
            if(yokoRank[x][i]>=orig){
                if(yokoRank[x][i]+diff>ans[x][i]){
                    ans[x][i] = yokoRank[x][i]+diff;
                    validateCol(x, i, yokoRank[x][i]+diff);
                }
            }
        }
    }
    private void validateCol(int x, int y, int after){
        //if(x==3&&y==8)  System.out.println("Call vCol at "+Integer.toString(x)+" "+Integer.toString(y)+" with "+Integer.toString(after));
        
        
        if(after<tatteRank[x][y]){
            ans[x][y] = tatteRank[x][y];
            validateRow(x, y, tatteRank[x][y]);
            //validateCol(x, y, tatteRank[x][y]);
            return;
        }
        /*
        if(after<yokoRank[x][y]){
            ans[x][y] = yokoRank[x][y];
            //validateRow(x, y, yokoRank[x][y]);
            validateCol(x, y, yokoRank[x][y]);
            return;
        }*/
        int orig = tatteRank[x][y];
        int diff = after - orig;
        for(int i = 0 ; i<m ; i++){
            if(i==x)    continue;
            if(tatteRank[i][y]>=orig){
                if(tatteRank[i][y]+diff>ans[i][y]){
                    ans[i][y] = tatteRank[i][y]+diff;
                    validateRow(i, y, tatteRank[i][y]+diff);
                }
            }
        }
    }
}