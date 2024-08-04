package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	var minLength, sum int
	l, r := 0, 0
	minLength = 1<<63 - 1
	if len(nums) == 0 {
		return 0
	}
	sum = nums[0]
	for r < len(nums) {
		if sum < target {
			r++
			sum += nums[r]
		} else {
			if r-l+1 < minLength {
				minLength = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}
	if minLength == 1<<63-1 {
		return 0
	}
	return minLength
}

func minSubArrayLen(target int, nums []int) int {
	minLength := 1<<63 - 1
	var l, sum int
	for r := range nums {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < minLength {
				minLength = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}
	if minLength == 1<<63-1 {
		return 0
	}
	return minLength
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}
