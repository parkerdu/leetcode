package main

func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	n := len(nums)
	if n == 1 {
		return true
	}
	dp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= n-1 {
			dp[i] = 1
			continue
		}
		for j := i + 1; j <= i+nums[i]; j++ {

			dp[i] |= dp[j]
		}
	}
	if dp[0] == 1 {
		return true
	}
	return false
}
