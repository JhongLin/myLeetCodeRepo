//https://leetcode.com/problems/domino-and-tromino-tiling/


func numTilings(n int) int {
    mod := 1000000007
    if n==1{
        return 1
    }else if n==2 {
        return 2
    }
    dp := make([]int, n+1)
    dp[1], dp[2] = 1, 2
    for i:=3 ; i<=n ; i++ {
        dp[i] += 2
        for j:=i-1 ; j>0 ; j-- {
            if j>2 {
                dp[i]+=2*dp[i-j]
            }else{
                dp[i]+=dp[i-j]
            }
            dp[i] %= mod
        }
    }
    return dp[n]
}