//https://leetcode.com/problems/partition-to-k-equal-sum-subsets/


class Solution {
    public boolean canPartitionKSubsets(int[] nums, int k) {
        int sum = 0;
        for( int value :nums )  sum+=value;
        if(sum%k!=0||k>nums.length)    return false;
        Arrays.sort(nums);
        return dfs(nums, 0, nums.length-1, sum/k, k);
    }
    private boolean dfs(int[] A, int sum, int start, int target, int k) {
        if(k==1)    return true;
        if(sum==target) return dfs(A, 0, A.length-1, target, k-1);
        for(int i = start ; i>=0 ; i--){
            if( A[i]!=-1 && sum+A[i]<=target ){
                int t = A[i];
                A[i] = -1;
                if(dfs(A, sum+t, i-1, target, k))    return true;
                A[i] = t;
            }
        }
        return false;
    }
}