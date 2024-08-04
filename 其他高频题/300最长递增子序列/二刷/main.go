package main

import "fmt"

// 18:00 - 18:33 还是在二分查找的代码上不熟悉
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tail := make([]int, 0)
	length := 0
	for i := range nums {
		index := findIndex(tail, nums[i])
		if index == length {
			tail = append(tail, nums[i])
			length++
		} else {
			tail[index] = nums[i]
		}
	}
	return length
}

// 在排序数组中找 比target所在位置, 没有比他小的返回0
func findIndex(tail []int, target int) int {
	if len(tail) == 0 || target < tail[0] {
		return 0
	}
	l, r := 0, len(tail)-1
	for l <= r {
		mid := l + (r-l)/2
		if tail[mid] == target {
			return mid
		} else if tail[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	fmt.Println(findIndex([]int{2, 4, 7}, 8))
}
