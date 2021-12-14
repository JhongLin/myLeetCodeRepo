//https://leetcode.com/problems/partition-equal-subset-sum/


func canPartition(nums []int) bool {
    n:=len(nums)
    if n==1 {return false}
    sum:=0
    for _,val := range nums{
        sum+=val
    }
    if sum%2==1 {return false}
    sum/=2
    dp := make([]bool, sum+1)
    dp[0] = true
    for _,val := range nums {
        for i:=sum ; i>=val ; i-- {
            dp[i] = dp[i]||dp[i-val]
        }
    }
    return dp[sum]
}