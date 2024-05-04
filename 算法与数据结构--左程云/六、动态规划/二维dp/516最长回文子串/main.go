package main

/*
题目4
最长回文子序列
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度
测试链接 : https://leetcode.cn/problems/longest-palindromic-subsequence/
 */
// 法1： 使用一行做dp, 从最后一行开始向上更新，每行先更新j=i,i+1, 再运用转移函数更新j=i+2,i+3...n-1
// 时间0(n**2), 空间O(N)
func longestPalindromeSubseq(s string) int {
	// 定义i行的dp[j] 表示s[i:j]的最长子序列(闭区间包括j)
	n := len(s)
	dp := make([]int,n)
	// 从最后一行开始，列从i+2开始
	for i := n-1; i>=0; i--{
		dp[i] = 1
		leftDown := 0
		if j := i+1; j < n  {
			leftDown = dp[j]
			if s[i] == s[j] {
				dp[j]= 2
			}
		}
		for j := i+2; j < n; j++ {
			backup := dp[j]   // 先备份一下，最后返回时只为leftDown
			if s[i] == s[j] {
				dp[j] = 2 + leftDown
			} else {
				ans := dp[j]
				if dp[j-1] > ans {
					ans = dp[j-1]
				}
				dp[j] = ans
			}
			leftDown = backup
		}
	}
	return dp[n-1]
}

// 法二： 使用一列，从左向右更新，每一列先更新i=j,j-1, 再运用转移函数更新i=j-2
func longestPalindromeSubseq1(s string) int {
	// 定义i行的dp[j] 表示s[i:j]的最长子序列(闭区间包括j)
	n := len(s)
	dp := make([]int,n)
	// 从最后一行开始，列从i+2开始
	for i := n-1; i>=0; i--{
		dp[i] = 1
		leftDown := 0
		if j := i+1; j < n  {
			leftDown = dp[j]
			if s[i] == s[j] {
				dp[j]= 2
			}
		}
		for j := i+2; j < n; j++ {
			backup := dp[j]   // 先备份一下，最后返回时只为leftDown
			if s[i] == s[j] {
				dp[j] = 2 + leftDown
			} else {
				ans := dp[j]
				if dp[j-1] > ans {
					ans = dp[j-1]
				}
				dp[j] = ans
			}
			leftDown = backup
		}
	}
	return dp[n-1]
}