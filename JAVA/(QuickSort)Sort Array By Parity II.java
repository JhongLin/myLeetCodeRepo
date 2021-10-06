//https://leetcode.com/problems/sort-array-by-parity-ii/


class Solution {
    public int[] sortArrayByParityII(int[] nums) {
        int position=0;
        int i = 1;
        while(true){
            while(position<nums.length&&nums[position]%2==0) position+=2;
            while(i<nums.length&&nums[i]%2==1)  i+=2;
            if(position<nums.length&&i<nums.length)    swap(nums, position, i);
            else break;
        }
        //System.out.println(Arrays.toString(nums));
        if(nums.length>2){
            mySort(nums, 0, nums.length-2);
            mySort(nums, 1, nums.length-1);
        }
        return nums;
    }
    private void mySort(int[] arr, int start, int end){
        int pivot = arr[end];
        int j = start-2;
        for(int i = start ; i<end ; i+=2){
            if(arr[i]<pivot){
                j+=2;
                swap(arr, i, j);
            }
        }
        j+=2;
        swap(arr, end, j);
        if(j>start+2)   mySort(arr, start, j-2);
        if(j<end-2) mySort(arr, j+2, end);
    }

    
    private void swap(int[] arr, int a, int b){
        int temp = arr[a];
        arr[a] = arr[b];
        arr[b] = temp;
    }
}