class Solution {
    public int partitionDisjoint(int[] nums) {
        /*
        int i = len/2;
        int[] left = Arrays.copyOf(nums, i);
        int[] right = Arrays.copyOfRange(nums, i, len);
        System.out.println(Arrays.toString(left));
        System.out.println(Arrays.toString(right));
        return i;
        */
        int pMax = -1;
        int right = nums.length-1;
        //int prevLen = 1;
        for(int i = 0 ; i<right ; i++){
            if(pMax<nums[i]){
                pMax = nums[i];
                for(int j = nums.length-1 ; j>=0 ; j--) {
                    if(j<=i)    return i+1;
                    else if(nums[j]<pMax){
                        right = j;
                        //i = j;
                        break;
                    }
                }
            }
        }
        //System.out.println(Arrays.toString(pMax));
        return right+1;
    }
}