# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        res = 0
        frontMin = prices[0]
        for i in range(1, len(prices)):
            if prices[i]-frontMin>res:
                res = prices[i]-frontMin
            if prices[i]<frontMin:
                frontMin = prices[i]
        return res