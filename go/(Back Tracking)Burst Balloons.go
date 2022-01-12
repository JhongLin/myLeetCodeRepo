//https://leetcode.com/problems/burst-balloons/

func maxCoins(nums []int) int {
    n := 1
    exNums := make([]int, len(nums)+2)
    for _, val := range nums {
        if val>0 {
            exNums[n] = val
            n++
        }
    }
    exNums[0], exNums[n] = 1 , 1
    n++
    memo := make([][]int, n)
    for i:=0 ; i<n ; i++ {memo[i]=make([]int, n)}
    return burst(memo, exNums, 0, n-1)
}

func burst(memo [][]int, nums []int, left, right int)int {
    if left+1==right {return 0}
    if memo[left][right]>0 {return memo[left][right]}
    res := 0
    for i:=left+1 ; i<right ; i++ {
        res = max(res, nums[left]*nums[i]*nums[right]+burst(memo, nums, left, i)+burst(memo, nums, i, right))
    }
    memo[left][right] = res
    return res
}

func max(a, b int) int {
    if a<b {
        return b
    }else{
        return a
    }
}