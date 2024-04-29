class Solution:
    def maxProfit(self, prices) -> int:
        for i in range(len(prices)-1):
            result = 0
            if prices[i+1] <= prices[i]:
                r = 0
            else:
                r = prices[i+1] - prices[i]
            result += r
        return result

