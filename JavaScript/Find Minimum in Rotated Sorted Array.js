/**
 * @param {number[]} nums
 * @return {number}
 */
 var findMin = function(nums) {
    if(nums.length==1)  return nums[0];
    
    let index = Math.floor(nums.length/2);
    let done = false;
    let left = 0;
    let right = nums.length-1;
    while(!done){
        //console.log(index);
        let l = index==0? nums.length-1:index-1;
        let r = index==nums.length-1? 0:index+1;
        if(nums[l]<nums[index]&&nums[index]<nums[r]){
            //console.log(nums[left]+' '+nums[index]+' '+nums[right]);
            if(nums[left]<nums[index]&&nums[index]<nums[right]){
                //console.log("O");
                right = index-1;
                index = Math.floor((left+right)/2);
            }else if(nums[left]<nums[index]&&nums[index]>nums[right]){
                //console.log("OOO");
                if(nums[right]<nums[left]){
                    left = index+1;
                    index = Math.floor((left+right)/2);
                }else{
                    right = index-1;
                    index = Math.floor((left+right)/2);
                }
            }else{
                //console.log("OO");
                if(nums[right]>nums[left]){
                    left = index+1;
                    index = Math.floor((left+right)/2);
                }else{
                    right = index-1;
                    index = Math.floor((left+right)/2);
                }
            }
        }else{
            return Math.min(nums[index], nums[r]);
        }
        //break;
    }
    return nums[index];
};