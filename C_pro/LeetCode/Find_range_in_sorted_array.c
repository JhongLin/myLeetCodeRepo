// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/



/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int findPosition(int* nums, int left, int right, int target)
{
    int mid = (left+right) / 2;
    while(left<right)
    {
        if(nums[mid]==target)
        {
            return mid;
        }
        else if(nums[mid]<target)
        {
            left = mid+1;
            mid = (left+right) / 2;
        }
        else
        {
            right = mid-1;
            mid = (left+right) / 2;
        }
    }
    if(left==right)
    {
        if(nums[left]==target)  return left;
        else    return -1;
    }
    return -1;
}


int* searchRange(int* nums, int numsSize, int target, int* returnSize){
    int *ans = (int*)malloc(2*sizeof(int));
    *returnSize = 2;
    ans[0] = -1, ans[1] = -1;
    if(numsSize==0) return ans;
    
    int pos = findPosition(nums, 0, numsSize-1, target);
    if(pos==-1) return ans;
    int left=pos, right=pos;
    
    while(left>=0)
    {
        
        ans[0] = left;
        left = findPosition(nums, 0, left-1, target);
        //printf("%d ", left);
    }
    //puts("");
    while(right<numsSize&&right>=0)
    {
        //printf("O");
        ans[1] = right;
        right = findPosition(nums, right+1, numsSize-1, target);
        //printf("%d ", right);
    }
    return ans;
}