class Solution {
    public int threeSumClosest(int[] nums, int target) {
        int ans = nums[0]+nums[1]+nums[2];
        Arrays.sort(nums);
        //System.out.println(Arrays.toString(nums));
        for(int i = 0 ; i<nums.length-2 ; i++){
            int start = i+1, end = nums.length-1;
            
            while(start<end){
                int sum = nums[i] + nums[start] + nums[end];
                if(Math.abs(sum-target)<Math.abs(ans-target))   ans = sum;
                if(sum>target)  end--;
                else if(sum<target) start++;
                else    return sum;
            }
        }
        return ans;
    }
}