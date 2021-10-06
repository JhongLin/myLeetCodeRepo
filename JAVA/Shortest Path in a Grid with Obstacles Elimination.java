//https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/


public class node {
    public int x;
    public int y;
    public int k;
    public node(int a, int b, int j){
        this.x = a;
        this.y = b;
        this.k = j;
    }
}

class Solution {
    public int shortestPath(int[][] grid, int k) {

        int row = grid.length;
        int col = grid[0].length;
        if(row==1&&col==1)  return 0;
        boolean[][] walked = new boolean[row][col];
        for(int i = 0 ; i<row ; i++)    Arrays.fill(walked[i], false);
        walked[0][0] = true;
        int[][] leftK = new int[row][col];
        leftK[0][0] = k;
        
        Queue<node> myQueue = new LinkedList<>();
        myQueue.offer(new node(0,0,k));
        int layerCount = 1;
        int feet = 0;
        boolean Done = false;
        
        
        while(layerCount>0&&!Done){
            int nextLayerCount=0;
            feet++;
            for(int n = 0 ; n<layerCount ; n++){
                node cur = myQueue.poll();
                if(cur.x+1<row&&!(walked[cur.x+1][cur.y])){
                    if(grid[cur.x+1][cur.y]==1&&cur.k>0){
                        if(cur.x+1==row-1&&cur.y==col-1){
                            Done = true;
                            break;
                        }
                        walked[cur.x+1][cur.y] = true;
                        leftK[cur.x+1][cur.y] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x+1,cur.y,cur.k-1));
                    }else if(grid[cur.x+1][cur.y]==0){
                        if(cur.x+1==row-1&&cur.y==col-1){
                            Done = true;
                            break;
                        }
                        walked[cur.x+1][cur.y] = true;
                        leftK[cur.x+1][cur.y] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x+1,cur.y,cur.k));
                    }
                }else if(cur.x+1<row&&(walked[cur.x+1][cur.y])){
                    if(grid[cur.x+1][cur.y]==1&&cur.k-1>leftK[cur.x+1][cur.y]){
                        if(cur.x+1==row-1&&cur.y==col-1){
                            Done = true;
                            break;
                        }
                        leftK[cur.x+1][cur.y] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x+1,cur.y,cur.k-1));
                        //System.out.println("O");
                    }else if(grid[cur.x+1][cur.y]==0&&cur.k>leftK[cur.x+1][cur.y]){
                        if(cur.x+1==row-1&&cur.y==col-1){
                            Done = true;
                            break;
                        }
                        leftK[cur.x+1][cur.y] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x+1,cur.y,cur.k));
                        //System.out.println("OO");
                    }
                }
                
                
                if(cur.y+1<col&&!(walked[cur.x][cur.y+1])){
                    if(grid[cur.x][cur.y+1]==1&&cur.k>0){
                        if(cur.x==row-1&&cur.y+1==col-1){
                            Done = true;
                            break;
                        }
                        walked[cur.x][cur.y+1] = true;
                        leftK[cur.x][cur.y+1] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y+1,cur.k-1));
                    }else if(grid[cur.x][cur.y+1]==0){
                        if(cur.x==row-1&&cur.y+1==col-1){
                            Done = true;
                            break;
                        }
                        walked[cur.x][cur.y+1] = true;
                        leftK[cur.x][cur.y+1] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y+1,cur.k));
                    }
                }else if(cur.y+1<col&&(walked[cur.x][cur.y+1])){
                    if(grid[cur.x][cur.y+1]==1&&cur.k-1>leftK[cur.x][cur.y+1]){
                        if(cur.x==row-1&&cur.y+1==col-1){
                            Done = true;
                            break;
                        }
                        leftK[cur.x][cur.y+1] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y+1,cur.k-1));
                        //System.out.println("X");
                    }else if(grid[cur.x][cur.y+1]==0&&cur.k>leftK[cur.x][cur.y+1]){
                        if(cur.x==row-1&&cur.y+1==col-1){
                            Done = true;
                            break;
                        }
                        leftK[cur.x][cur.y+1] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y+1,cur.k));
                        //System.out.println("XX");
                    }
                }
                
                
                
                if(cur.x-1>=0&&!(walked[cur.x-1][cur.y])){
                    if(grid[cur.x-1][cur.y]==1&&cur.k>0){

                        walked[cur.x-1][cur.y] = true;
                        leftK[cur.x-1][cur.y] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x-1,cur.y,cur.k-1));
                    }else if(grid[cur.x-1][cur.y]==0){

                        walked[cur.x-1][cur.y] = true;
                        leftK[cur.x-1][cur.y] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x-1,cur.y,cur.k));
                    }
                }else if(cur.x-1>=0&&(walked[cur.x-1][cur.y])){
                    if(grid[cur.x-1][cur.y]==1&&cur.k-1>leftK[cur.x-1][cur.y]){
                        leftK[cur.x-1][cur.y] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x-1,cur.y,cur.k-1));
                    }else if(grid[cur.x-1][cur.y]==0&&cur.k>leftK[cur.x-1][cur.y]){
                        leftK[cur.x-1][cur.y] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x-1,cur.y,cur.k));
                    }
                }
                
                if(cur.y-1>=0&&!(walked[cur.x][cur.y-1])){
                    if(grid[cur.x][cur.y-1]==1&&cur.k>0){

                        walked[cur.x][cur.y-1] = true;
                        leftK[cur.x][cur.y-1] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y-1,cur.k-1));
                    }else if(grid[cur.x][cur.y-1]==0){

                        walked[cur.x][cur.y-1] = true;
                        leftK[cur.x][cur.y-1] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y-1,cur.k));
                    }
                }else if(cur.y-1>=0&&(walked[cur.x][cur.y-1])){
                    if(grid[cur.x][cur.y-1]==1&&cur.k-1>leftK[cur.x][cur.y-1]){
                        leftK[cur.x][cur.y-1] = cur.k-1;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y-1,cur.k-1));
                    }else if(grid[cur.x][cur.y-1]==0&&cur.k>leftK[cur.x][cur.y-1]){
                        leftK[cur.x][cur.y-1] = cur.k;
                        nextLayerCount++;
                        myQueue.offer(new node(cur.x,cur.y-1,cur.k));
                    }
                }
                
            }// end for (this layer)
            layerCount = nextLayerCount;
        }
        //for(int i = 0; i<row ; i++) System.out.println(Arrays.toString(grid[i]));
        if(!Done)   return -1;
        return feet;
    }
}