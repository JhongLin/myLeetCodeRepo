//https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
    res := nums[0];
    curr := nums[0];
    for i := 1; i < len(nums) ; i++ {
        if curr < 0{
            curr = nums[i]
        }else{
            curr +=nums[i]
        }
        if res < curr {
            res = curr;
        }
    }
    return res
}
