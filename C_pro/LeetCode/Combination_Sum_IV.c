//https://leetcode.com/explore/item/3713

void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}


void quickSort(int *nums, int start, int end)
{
    int pivot = nums[end], smaller = start-1;
    for(int i = start ; i<end ; i++)
    {
        if(nums[i]<=pivot) swap(&nums[++smaller], &nums[i]);
    }
    swap(&nums[++smaller], &nums[end]);
    if(start<smaller-1) quickSort(nums, start, smaller-1);
    if(end>smaller+1) quickSort(nums, smaller+1, end);
}

void treeRecur(int *nums, int numsSize, int left, int *ans)
{   
    
    if(left==0)
    {
        //printf("O");
        *ans += 1;
        return;
    }
    for(int i = 0; i<numsSize; i++)
    {
        if(nums[i]<=left) treeRecur(nums, numsSize, left-nums[i], ans);
        else return;
    }
}

int combinationSum4(int* nums, int numsSize, int target){
    /*
    int ans=0;
    quickSort(nums, 0, numsSize-1);
    treeRecur(nums, numsSize, target, &ans);
    return ans;
    */
    unsigned int *dp = (unsigned int*)calloc(target+1, sizeof(unsigned int));
    dp[0] = 1;
    for(int i = 1; i<=target ; i++)
        for(int j=0; j<numsSize ; j++)
            if(nums[j]<=i) dp[i] += dp[i-nums[j]];
    return (int)dp[target];
}