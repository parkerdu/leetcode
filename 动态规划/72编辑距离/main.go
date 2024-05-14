package main

import "fmt"

// 18:12 - 19:02 部分case还有问题，空间压缩不熟悉, 插入操作没考虑两个问题导致
func minDistance(word1 string, word2 string) int {
	// step1: 定义dp[i][j]为长度为i的word1和长度为j的word2的编辑距离
	m, n := len(word1), len(word2)
	// step2: 初始化一行数据，长度为n+1, 元素为0-n
	dp := make([]int, 1+n)
	for i := range dp {
		dp[i] = i
	}
	// step3: 从第1行开始做dp,
	var leftUp int
	for i := 1; i <= m; i++ {
		leftUp = dp[0]
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = leftUp
			} else {
				min := 1 + leftUp
				if 1+dp[j-1] < min {
					min = 1 + dp[j-1]
				}
				if 1+dp[j] < min {
					min = 1 + dp[j]
				}
				dp[j] = min
			}
			leftUp = tmp
		}
		// 最后更新dp[0]
		dp[0] = i
	}
	// step4: 最终返回dp[m][n]
	return dp[n]
}

func main() {
	s1 := "horse"
	s2 := "ros"
	fmt.Println(minDistance(s1, s2))
}
