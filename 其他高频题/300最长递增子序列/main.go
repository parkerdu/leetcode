package main

import "fmt"

func lengthOfLIS(nums []int) int {
	tails := make([]int, 0)
	length := 0
	for i := range nums {
		//
		index := findIndex1(tails, nums[i])
		if index == length {
			tails = append(tails, nums[i])
			length++
		} else {
			tails[index] = nums[i]
		}
	}
	return length
}

// 查找当前target在tails递增数组中的位置，存在返回下标，不存在返回比target严格小的最大值
// 找不到比target小的则返回0
// []左闭，右闭
func findIndex(tails []int, target int) int {
	l, r := 0, len(tails)-1
	for l <= r {
		mid := (l + r) / 2
		if tails[mid] < target {
			l = mid + 1
		} else if tails[mid] == target {
			return mid
		} else {
			r = mid - 1
		}
	}
	// 跳出for说明没找到，此时l=r+1
	return l
}

// []左闭，右开
func findIndex1(tails []int, target int) int {
	l, r := 0, len(tails)
	for l < r {
		mid := (l + r) / 2
		if tails[mid] < target {
			l = mid + 1
		} else if tails[mid] == target {
			return mid
		} else {
			r = mid
		}
	}
	// 跳出for说明没找到，此时l=r+1
	return l
}

func main() {
	t := []int{2, 3, 5}
	fmt.Println(findIndex(t, 1))
}
