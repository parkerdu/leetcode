package main

func canJump(nums []int) bool {
	// 定义变量curIndex = n-1
	// step1: 从n-2开始向前遍历，若index+value >= curIndex 则继续向前，更新curIndex为当前index
	n := len(nums)
	if n == 1 {
		return true
	}
	curIndex := n-1
	for j := n-2; j>=0; j-- {
		if j + nums[j] >= curIndex {
			curIndex = j
		}
	}
	if nums[0] < curIndex {
		return false
	}
	return true
}