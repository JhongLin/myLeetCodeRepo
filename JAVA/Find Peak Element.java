class Solution {
    public int findPeakElement(int[] nums) {
        int len=nums.length;
        int rear = 0, front = len - 1, mid;
        if(len==1)  return 0;
        if (nums[front]>nums[front-1])  return front;
        else if(nums[rear]>nums[rear+1])    return rear;
        mid = (front+rear)/2;
        while(true){
            if(nums[mid+1]<nums[mid]){
                if(nums[mid-1]<nums[mid])   return mid;
                else{
                    front = mid;
                    mid = (front+rear)/2;
                }
            }else {
                rear = mid;
                mid = (front+rear)/2;
            }
            
        }
        
        //return 0;
    }
}