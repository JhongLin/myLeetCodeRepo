//https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/

int numSubmatrixSumTarget(int** matrix, int matrixSize, int* matrixColSize, int target){
    int ans = 0;
    
    
    int **prefixSum = (int**)malloc(matrixSize*sizeof(int*));
    for(int i = 0 ; i<matrixSize ; i++)
    {
        prefixSum[i] = (int*)calloc(*matrixColSize, sizeof(int));
        prefixSum[i][0] = matrix[i][0];
    }
    for(int i = 0 ; i<matrixSize ; i++)
    {
        for(int j = 1 ; j<*matrixColSize ; j++)
        {
            prefixSum[i][j] = prefixSum[i][j-1]+matrix[i][j];
        }
    }
    for(int i = 1 ; i<matrixSize ; i++)
    {
        for(int j = 0 ; j<*matrixColSize ; j++)
        {
            prefixSum[i][j] += prefixSum[i-1][j];
        }
    }

    for(int y = 1 ; y<=matrixSize ; y++ )
    {
        for(int x = 1 ; x<=*matrixColSize ; x++ )
        {
            for(int i = y-1 ; i<matrixSize ; i++) //row number
            {
                for(int j = x-1 ; j<*matrixColSize ; j++) // column number
                {   
                    int leftSide=0, upSide=0, coverSide=0;
                    if(j-x>=0) leftSide = prefixSum[i][j-x];
                    if(i-y>=0) upSide = prefixSum[i-y][j];
                    if(j-x>=0&&i-y>=0) coverSide = prefixSum[i-y][j-x];
                    int val = prefixSum[i][j] -leftSide -upSide +coverSide;
                    if (val == target)
                    {                        
                        ans++;
                        //printf("%d %d %d %d\n", y, x, i, j);
                    }
                    
                }
            }
        }
    }
    
    

    return ans;
}