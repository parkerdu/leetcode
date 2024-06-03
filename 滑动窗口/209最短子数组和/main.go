package main

func minSubArrayLen(target int, nums []int) int {
	var l, sum int
	res := 1<<63 - 1
	for r := range nums {
		sum += nums[r]
		for sum >= target {
			// 满足条件就结算
			if r-l+1 < res {
				res = r - l + 1
			}
			sum = sum - nums[l]
			l++
		}
		// l不能走了且满足条件结算
	}
	if res == 1<<63-1 {
		return 0
	}
	return res
}
