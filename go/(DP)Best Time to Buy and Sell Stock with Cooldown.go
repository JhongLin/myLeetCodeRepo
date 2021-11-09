//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/


func maxProfit(prices []int) int {
    n := len(prices)
    dp := make( []int , n+1 );
    min := prices[0]
    for i:=1 ; i<n+1 ; i++ {
        if i>=2 {
            if min > prices[i-1]-dp[i-2] {
                min = prices[i-1]-dp[i-2]
            }
        }
        dp[i] = Max(dp[i-1], prices[i-1]-min)
    }
    return dp[n]
}

func Max(x, y int) int {
    if x < y {
        return y
    }else{
        return x
    }
}