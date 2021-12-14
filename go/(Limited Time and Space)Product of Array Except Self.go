//https://leetcode.com/problems/product-of-array-except-self/


func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    ans[0] = 1
    for i:=1 ; i<n ; i++ {
        ans[i] = ans[i-1]*nums[i-1]
    }
    rightBack := 1
    for i:=n-1 ; i>=0 ; i-- {
        ans[i] *= rightBack
        rightBack *= nums[i]
    }
    return ans
}