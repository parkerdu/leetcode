package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return process(nums, 0, len(nums)-1).msum
}

// 在[l,r]闭区间上，
/*
 求包含l的最大子数组和lsum,
   包含r的最大子数组和rsum
   l,r区间内的最大子数组和，可能包含l,r，可能不包含
   区间和sum
*/

type State struct {
	lsum int
	rsum int
	msum int
	sum  int
}

func process(nums []int, l, r int) State {
	if l == r {
		return State{nums[l], nums[l], nums[l], nums[l]}
	}
	mid := l + (r-l)/2
	ls := process(nums, l, mid)
	rs := process(nums, mid+1, r)
	s := State{}
	s.lsum = max(ls.lsum, ls.sum+rs.lsum)
	s.rsum = max(rs.rsum, rs.sum+ls.rsum)
	s.sum = ls.sum + rs.sum
	s.msum = max(ls.msum, rs.msum, s.lsum, s.rsum, s.sum, ls.rsum+rs.lsum)
	return s
}
