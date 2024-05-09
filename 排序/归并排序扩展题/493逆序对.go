package main

import "fmt"

// 2024/05/09 23:13开始 -
/*
这个题和小和问题很像，可以先做./局部解决全部/小和问题.main
时间复杂度：O(N*logN)
空间：O(N) ?? mergeSORT空间是多少 对的，就是O(n）
关键是怎么把它和merge sort的merge结合起来
那就是左右在merge的过程中，由于都是排序的，所以若左边nums[i] > 2*nums[j] 则i右边的所有数都是满足这个条件的
*/
func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	process(nums, 0, len(nums)-1, &res)
	return res
}

func process(nums []int, left, right int, res *int) {
	if left == right {
		return
	}
	// step1: 找nums中点mid
	mid := left + (right-left)/2
	// step2: [left,mid]左边递归，
	process(nums, left, mid, res)
	process(nums, mid+1, right, res)
	// step3: 右边
	// step4: 合并，计算有几对
	merge(nums, left, mid, right, res)
}

func merge(nums []int, left, mid, right int, res *int) {
	// step1: 左右都不越界，比较左右的值，
	i, j := left, mid+1
	// 先compute，再做合并
	for i <= mid && j <= right {
		if nums[i] <= 2*nums[j] {
			i++
		} else {
			*res += mid - i + 1
			j++
		}
	}
	i, j = left, mid+1
	help := make([]int, right-left+1)
	index := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			help[index] = nums[i]
			i++
		} else {
			help[index] = nums[j]
			j++
		}
		index++
	}
	for i <= mid {
		help[index] = nums[i]
		i++
		index++
	}
	for j <= right {
		help[index] = nums[j]
		j++
		index++
	}
	// 复制
	for i := range help {
		nums[i+left] = help[i]
	}
}

func main() {
	l := []int{5, 4, 3, 2, 1}
	fmt.Println(reversePairs(l))
}
