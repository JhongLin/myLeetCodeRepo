//https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
    n := len(nums)
    toru := make([]int , n)
    toranai := make([]int, n)
    toru[0] = nums[0]
    for i:=1 ; i<n ; i++ {
        toru[i] = nums[i] + toranai[i-1]
        toranai[i] = max(toru[i-1], toranai[i-1])
    }
    
    
    return max(toranai[n-1], toru[n-1])
}

func max (a, b int) int {
    if a>b {
        return a
    }else{
        return b
    }
}