//https://leetcode.com/problems/maximum-product-subarray/

func maxProduct(nums []int) int {
    res, posMax, negMax := nums[0], nums[0], nums[0]
    n := len(nums)
    for i:=1 ; i<n ; i++ {
        temp := posMax
        posMax = max(nums[i], posMax*nums[i])
        posMax = max(posMax, negMax*nums[i])
        negMax = min(nums[i], negMax*nums[i])
        negMax = min(negMax, temp*nums[i])
        res = max(res, posMax)
    }
    
    return res
}

func max (a, b int) int {
    if a>b {
        return a
    }else{
        return b
    }
}
func min (a, b int) int {
    if a>b {
        return b
    }else{
        return a
    }
}