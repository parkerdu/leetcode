package main

import "math"

/*
思路：递归求解，
 */
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	p, _ := f1(prices, len(prices)-1)
	return p
}

// 定义f1为求prices[0:i+1]的最大利润, 返回当前最大利润和当前的有效最大值
func f1(p []int, i int) (int, int) {
	if i == 0 {
		return 0, p[i]
	}
	profit, min := f1(p,i-1)
	if p[i] - min > profit {
		return p[i]-min, min
	}
	if p[i] < min {
		min = p[i]
	}
	return profit, min
}

// 法二：从左向右遍历，每次记住当前的最小值
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var min,res int
	min = math.MaxInt64
	for i, _ := range prices {
		if prices[i] > min {
			p := prices[i] - min
			if p > res {
				res = p
			}
		} else {
			min = prices[i]
		}
	}
	return res
}
