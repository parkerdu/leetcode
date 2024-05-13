package main

import "fmt"

/*
思路：堆排序 大顶堆，弹出k-1个，当前的root节点为答案
法二：维护一个K size的小根堆，nums中比堆顶大的进来，小的跳过，最终答案为堆顶
*/
func findKthLargest(nums []int, k int) int {
	// step1: 前k个元素入堆，构造出k-size小根堆，headInsert
	// step2: 比堆顶大，则交互并heapify
	for i := range nums {
		if i < k {
			heapInsert(nums, i)
		} else if nums[i] > nums[0] {
			swap(nums, 0, i)
			heapify(nums, 0, k)
		}
	}
	return nums[0]
}

// 小根堆插入，比父亲小往上
func heapInsert(nums []int, index int) {
	for nums[index] < nums[(index-1)/2] {
		swap(nums, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// 当前index开始，向下堆化，比孩子大向下
func heapify(nums []int, index, heapSize int) {
	for 2*index+1 < heapSize {
		min := 2*index + 1
		// 若右孩子存在，找最小值
		if 2*index+2 < heapSize && nums[2*index+2] < nums[min] {
			min = 2*index + 2
		}
		if nums[index] > nums[min] {
			swap(nums, index, min)
			index = min
		} else {
			return
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{1, 3, 456, 2, 7}
	fmt.Println(findKthLargest(nums, 2))
}
