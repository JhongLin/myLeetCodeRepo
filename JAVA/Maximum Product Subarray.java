//https://leetcode.com/problems/maximum-product-subarray/


class Solution {
    public int maxProduct(int[] nums) {
        int pMax=nums[0], nMax=nums[0], max = nums[0];
        for(int i = 1 ; i<nums.length ; i++){
            int temp = pMax;
            pMax = Math.max(nums[i], pMax*nums[i]);
            pMax = Math.max(pMax, nMax*nums[i]);
            
            nMax = Math.min(nMax*nums[i], nums[i]);
            nMax = Math.min(nMax, temp*nums[i]);
            if(max<pMax)    max = pMax;
        }
        return max;
    }
}