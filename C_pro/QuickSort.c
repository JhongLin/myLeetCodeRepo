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