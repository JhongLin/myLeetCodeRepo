//https://leetcode.com/problems/ones-and-zeroes/

int max(int a, int b)
{
    return a>b? a:b;
}

int findMaxForm(char ** strs, int strsSize, int m, int n){
    if(strs==NULL)
    {
        return 0;
    }
    int **dpTable = (int**)malloc((m+1)*sizeof(int*));
    for(int i = 0 ; i < m+1 ; i++ )
    {
        dpTable[i] = (int*)calloc(n+1, sizeof(int));
    }
    
    for(int i = 0 ; i<strsSize ; i++) //for each string do once
    {   
        int *weight = (int*)calloc(2, sizeof(int));
        //First count the 1s and 0s for each String
        for(int j = 0 ; j < strlen(strs[i]) ; j++ )
        {
            weight[strs[i][j]-'0']++;
        }
        for(int x = m ; x>=weight[0] ; x--)
        {
            for(int y = n ; y>=weight[1] ; y--)
            {
                dpTable[x][y] = max(dpTable[x][y], dpTable[x-weight[0]][y-weight[1]]+1);
            }
        }
        
        
        
        free(weight);
    }

    return dpTable[m][n];
}