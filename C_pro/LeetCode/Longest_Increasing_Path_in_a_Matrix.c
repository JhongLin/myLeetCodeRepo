//https://leetcode.com/problems/longest-increasing-path-in-a-matrix/


int dfs(int **matrix, int matrixSize, int matrixColSize, int x, int y, int **mp)
{   //printf("DFS now handling %d %d\n", x, y);
    int arrows[4]={};
    if(x-1>=0)//to upside
    {
        if(mp[x-1][y]&&matrix[x-1][y]>matrix[x][y]) arrows[0]=1+mp[x-1][y];
        else if(matrix[x-1][y]>matrix[x][y]) arrows[0]=1+dfs(matrix, matrixSize, matrixColSize, x-1, y, mp);
    }
    if(x+1<matrixSize)//to downside
    {
        if(mp[x+1][y]&&matrix[x+1][y]>matrix[x][y]) arrows[1]=1+mp[x+1][y];
        else if(matrix[x+1][y]>matrix[x][y]) arrows[1]=1+dfs(matrix, matrixSize, matrixColSize, x+1, y, mp);
    }
    if(y-1>=0) //to leftside
    {
        if(mp[x][y-1]&&matrix[x][y-1]>matrix[x][y]) arrows[2]=1+mp[x][y-1];
        else if(matrix[x][y-1]>matrix[x][y]) arrows[2]=1+dfs(matrix, matrixSize, matrixColSize, x, y-1, mp);
    }
    if(y+1<matrixColSize)//to rightside
    {
        if(mp[x][y+1]&&matrix[x][y+1]>matrix[x][y]) arrows[3]=1+mp[x][y+1];
        else if(matrix[x][y+1]>matrix[x][y]) arrows[3]=1+dfs(matrix, matrixSize, matrixColSize, x, y+1, mp);
    }
    int maxArrow = 0;
    for(int i = 0 ; i<4 ; i++)
    {
        if(maxArrow<arrows[i]) maxArrow=arrows[i];
    }
    //printf("%d %d is done\n", x, y);
    if (maxArrow==0)
    {   
        //printf("no way for %d %d\n", x, y);
        mp[x][y] = 1;
        return 1;
    }
    else
    {
        mp[x][y] = maxArrow;
        return maxArrow;
    }
}


int longestIncreasingPath(int** matrix, int matrixSize, int *matrixColSize){
    /*
    if(matrixSize==1)
    {   int journey = 1, maxJourney = 1;
        for(int i = 0; i<*matrixColSize-1 ; i++)
        {
            if(matrix[0][i]<matrix[0][i+1])
            {
                journey++;
                if(maxJourney<journey) maxJourney=journey;
            }
            else
            {
                journey = 1;
            }
        }
        return maxJourney;
    }
    */
    
    int maxLen = 0;
    int **mp = (int**)malloc(matrixSize*sizeof(int*));
    for(int i = 0 ; i<matrixSize ; i++)
    {
        mp[i] = (int*)calloc(*matrixColSize, sizeof(int));
    }
    for(int i = 0 ; i < matrixSize ; i++ )
    {
        for(int j = 0 ; j < *matrixColSize ; j++ )
        {
            while(mp[i][j]==0)
            {
                //mp[i][j] = 1;
                dfs(matrix, matrixSize, *matrixColSize, i, j, mp);
            }
            if(mp[i][j]>maxLen) maxLen = mp[i][j];
        }
    }
    /*
    for(int i = 0 ; i < matrixSize ; i++ )
    {
        for(int j = 0 ; j < *matrixColSize ; j++ )
        {
            printf("%d ", mp[i][j]);
        }
        puts("");
    }*/
    return maxLen;
}