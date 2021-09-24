/**
 * @param {number[]} nums
 * @return {number}
 */
 var arrayNesting = function(nums) {
    let status = [];
    let max = 0;
    for(let i = 0; i<nums.length ; i++){
        if(status[i]===undefined){
            if(nums[i]==i)  status[i]=1;
            else    helper(nums, status, i);
            if(status[i]>max)   max = status[i];
        }
    }
    return max;
};
function helper(nums, arr, target){
    if(arr[nums[target]]===undefined){
        arr[target] = 0;
        helper(nums, arr, nums[target]);
        arr[target] = 1+arr[nums[target]];
    }else{
        arr[target] = 1+arr[nums[target]];
    }
}