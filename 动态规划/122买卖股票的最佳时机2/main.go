package main

/*
买卖股票的最佳时机2
可多次买卖

法一：只要是发生连续涨幅的我就把利润要了
*/

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res

}

/*
法二：dp
dp[i][0] 表示第i天持有股票时手上的最大现金
dp[i][1] 表示第i天不持有股票手上的最大现金
假设刚开始手上现金为0， 第一次买股票为赊账，那么最后一天手上的最大现金即为最大利润，所以直接返回dp[n-1][1]

状态转移：
dp[i][0] =  max{dp[i-1][0], dp[i-1][1]-prices[i]}
dp[i][1] = max{dp[i-1][1], dp[i-1][0]+prices[i]}

初始条件
dp[0][0] = -prices[0]
dp[0][1] = 0
*/
func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	for i, _ := range dp {
		dp[i] = [2]int{}
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
