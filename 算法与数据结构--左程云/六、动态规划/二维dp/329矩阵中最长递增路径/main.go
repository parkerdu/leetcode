package main

func longestIncreasingPath(matrix [][]int) int {
	// step1: 定义1个缓存数组dp, dp[i][j]表式m[i][j]开始的最长递增序列长度,初始化全为-1
	if len(matrix) == 0 {
		return 0
	}
	m,n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var ans int
	// step2: 按行，列，调用递归函数，并和ans比较，更新ans
	for i := range matrix {
		for j := range matrix[i] {
			if res := f1(dp, matrix, i,j);res > ans {
				ans = res
			}
		}
	}
	return ans
}

// f1(i,j) 求从m[i][j]开始的最长递增子序列
func f1(dp [][]int, m [][]int, i,j int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	next := 0
	if i-1>=0 && m[i-1][j] > m[i][j] {
		if res := f1(dp,m,i-1,j); res > next {
			next = res
		}
	}
	if i+1<len(m) && m[i+1][j] > m[i][j] {
		if res := f1(dp,m,i+1,j); res > next {
			next = res
		}
	}
	if j-1>=0 && m[i][j-1] > m[i][j] {
		if res := f1(dp,m,i,j-1); res > next {
			next = res
		}
	}
	if j+1 < len(m[i]) && m[i][j+1] > m[i][j] {
		if res := f1(dp,m,i,j+1); res > next {
			next = res
		}
	}
	dp[i][j] = 1+next
	return 1+next
}
