package main

import "math"

/*
动态规划
dp[i] = nums[i]到last的最少步数

if i + nums[i] >= n-1

	dp[i] =  1

else

	min{} + 1

初始条件：dp[n-1] = 0
*/
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= n-1 {
			dp[i] = 1
		} else {
			min := math.MaxInt - 100 // 应对1+min变为负数
			for j := i + 1; j <= i+nums[i]; j++ {
				if dp[j] < min {
					min = dp[j]
				}
			}
			dp[i] = 1 + min
		}
	}
	return dp[0]
}

/*
法二：我没想到，看评论的贪心算法
思路是：定义furthest  为下一次的最远点，curEnd为当前层的最后1个index, 遍历当前层更新further，当前层走完，说明需要跳一步
*/

func jump1(nums []int) int {
	var jumps, furthest, curEnd int
	if len(nums) == 1 {
		return 0
	}
	for i := range nums {
		furthest = max(furthest, i+nums[i])
		if furthest >= len(nums)-1 {
			jumps++
			break
		}
		if i == curEnd {
			jumps++
			curEnd = furthest
		}
	}
	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
