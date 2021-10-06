//https://leetcode.com/problems/partition-equal-subset-sum/


class Solution {
    public boolean canPartition(int[] nums) {
        int len = nums.length;
        if(len<2)   return false;
        int sum = 0;
        for(int val:nums)   sum+=val;
        if(sum%2==1)    return false;
        int target = sum/2;
        boolean[] dp = new boolean[target+1];
        dp[0] = true;
        for(int i = 0 ; i<len ; i++){
            for(int j = target; j>=nums[i]; j--){
                dp[j] = dp[j] || dp[j-nums[i]];
            }
        }
        return dp[target];
    }
}