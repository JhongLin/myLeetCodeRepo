//https://leetcode.com/problems/letter-combinations-of-a-phone-number/


const char *letters[] = { "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
const int strsSize[] = {0, 0, 3, 3, 3, 3, 3, 4, 3, 4};

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

char **declear2D(const int Dim1,const int Dim2){
    int i=0;
    char **p= (char **)malloc(Dim1 * sizeof(char **)); 

    for (i=0;i<Dim1;i++)
    { 
        *(p+i)=(char *)calloc(Dim2 , sizeof(char *)); 
        //for (int j=0; j<Dim2-1; j++) *(*(p+i)+j) = 'b'; 
        //*(*(p+i)+Dim2-1) = 0; 
    } 

    return p; 
}

char ** letterCombinations(char * digits, int* returnSize){
    
    int sLen=strlen(digits);
    
    if(sLen==0)
    {   
        *returnSize = 0;
        return NULL;
    }
    *returnSize = 1;
    int *serials = (int*)malloc(sLen*sizeof(int));
    for(int i = 0 ; i < sLen ; i++)
    {
        serials[i] = (int)digits[i]-48;
        *returnSize *= strsSize[serials[i]];
    }
    char **ans = declear2D(*returnSize, sLen);
    int interval = 1;
    for(int i = sLen-1 ; i>=0 ; i--)
    {
        for(int j = 0 ; j < strsSize[serials[i]] ; j++)
        {   
            for(int k = 0 ; k < *returnSize/(strsSize[serials[i]]*interval) ; k++)
            {
                for(int l = 0 ; l < interval ; l++)
                {
                    ans[k*strsSize[serials[i]]*interval+l+j*interval][i] = letters[serials[i]][j];
                }
            }
        }
        interval *= strsSize[serials[i]];
    }
    
    return ans;

}