//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
    hold := prices[0]
    topSell := prices[0]
    ans := 0
    for i:= 1 ; i<len(prices) ; i++ {
        if hold>=prices[i] {
            hold = prices[i]
            topSell = prices[i]
        }else{
            if topSell<prices[i] {
                ans += prices[i] - topSell
                topSell = prices[i]
            }else{
                if topSell > prices[i] {
                    hold = prices[i]
                    topSell = prices[i]
                }
            }
        }
    }
    return ans
}