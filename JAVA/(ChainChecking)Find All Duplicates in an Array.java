//https://leetcode.com/problems/find-all-duplicates-in-an-array/

class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        List<Integer> ans = new ArrayList<>();
        for(int i = 0 ; i<nums.length ; i++){
            if(nums[i]==0) continue;
            if(nums[i]!=i+1){
                if(nums[nums[i]-1]==nums[i]){
                    nums[nums[i]-1] = 0;
                    ans.add(nums[i]);
                    nums[i] = 0;
                }else{
                    while(nums[i]!=0&&nums[i]!=i+1){
                        if(nums[nums[i]-1]==nums[i]){
                            //System.out.println(Integer.toString(i)+' '+Integer.toString(nums[i]));
                            nums[nums[i]-1] = 0;
                            ans.add(nums[i]);
                            nums[i] = 0;
                        }else{
                            swap(nums, nums[i]-1, i);
                        }
                    }
                }
            }
        }
        return ans;
    }
    private void swap(int[] arr, int a, int b){
        int temp = arr[a];
        arr[a] = arr[b];
        arr[b] = temp;
    }
}