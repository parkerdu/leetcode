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

// 深信服一面：除了找出最小距离，还要打印变换的步骤，面试中为什么没写出来？因为dp[i][j]的定义弄错了，为a[:i]转换为b[:j]的距离，怎么表示空字符串？没办法表示，得换成从头开始的长度
func change(a,b string) int {
	// 定义dp[i][j] 为a[:i]转换为b[:j]的编辑距离， 数组元素为[2]string  前面表示操作，后面表示字符， m*n维度
	m,n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始条件
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := 1; i<=m; i++ {
		dp[i][0] = i
		for j :=1; j<=n; j++{
			if a[i-1] == b[j-1] {  // 相同
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 不相同，取增，删，改的最小者
				min := 1<<63-1
				if 1+dp[i][j] < min {  // 增加操作
					min = 1+dp[i][j-1]
				}
				if 1+dp[i-1][j-1] < min { // 替换
					min = 1 + dp[i-1][j-1]
				}
				if 1 + dp[i-1][j] < min {  // 删除当前i
					min = 1+ dp[i-1][j]
				}
				dp[i][j] = min
			}
		}
	}
	return dp[m][n]
}

func main() {
	s1 := "horse"
	s2 := "ros"
	fmt.Println(minDistance(s1, s2))
	fmt.Println(change(s1,s2))
}
