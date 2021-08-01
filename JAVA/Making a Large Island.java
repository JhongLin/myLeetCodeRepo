class Solution {
    public int nowSerial;
    public int nowArea;
    public int n;
    
    public int largestIsland(int[][] grid) {
        List<Integer> serial = new ArrayList<>();
        List<Integer> area = new ArrayList<>();
        n = grid.length;
        nowSerial = 0;
        for(int i = 0 ; i<n ; i++ ){
            for(int j=0;j<n;j++){
                if(grid[i][j]==1){
                    nowSerial--;
                    nowArea = 0;
                    countArea(grid, i, j);
                    serial.add(nowSerial);
                    area.add(nowArea);
                }
            }
        }
        if(serial.isEmpty())    return 1;
        else if(serial.size()==1){
            if(area.get(0)<n*n) return area.get(0)+1;
            return n*n;
        }
        
        int ans = 0;
        Integer[] arrayOfArea = area.toArray(new Integer[0]);
        //Integer[] element_rdv_id = ids_rdv.toArray(new Integer[0]);
        for(int i = 0 ; i<n ; i++ ){
            for(int j=0;j<n;j++){
                if(grid[i][j]==0){
                    int[] nei = new int[4];
                    if(i-1>=0&&grid[i-1][j]<0) nei[0] = grid[i-1][j];
                    if(i+1<n&&grid[i+1][j]<0) nei[1] = grid[i+1][j];
                    if(j-1>=0&&grid[i][j-1]<0) nei[2] = grid[i][j-1];
                    if(j+1<n&&grid[i][j+1]<0) nei[3] = grid[i][j+1];
                    Arrays.sort(nei);
                    int temp=0;
                    for(int k = 0; k<4 ; k++){
                        if(k!=3&&nei[k]==nei[k+1])    continue;
                        //if(nei[k]<0)    temp += area.get( nei[k]*(-1)-1 );
                        if(nei[k]<0)    temp += arrayOfArea[ nei[k]*(-1)-1 ];
                    }
                    if(temp>ans)    ans = temp;
                }
            }
        }
        
        /*
        System.out.print(Arrays.toString(serial.toArray()));
        System.out.println(' '+Arrays.toString(area.toArray()));
        
        for(int i = 0 ; i<n ; i++ ){
            for(int j=0;j<n;j++){
                System.out.print(Integer.toString(grid[i][j])+' ');
            }
            System.out.println("");
        }*/
        return ans+1;
    }
    private void countArea(int[][] grid, int a, int b){
        grid[a][b] = nowSerial;
        nowArea++;
        if(a-1>=0&&grid[a-1][b]==1) countArea(grid, a-1, b);
        if(a+1<n&&grid[a+1][b]==1) countArea(grid, a+1, b);
        if(b-1>=0&&grid[a][b-1]==1) countArea(grid, a, b-1);
        if(b+1<n&&grid[a][b+1]==1) countArea(grid, a, b+1);
    }
}