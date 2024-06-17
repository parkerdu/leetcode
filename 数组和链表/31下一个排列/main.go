package main

import (
	"fmt"
)

/*
法一：两次遍历查找
先从右向左遍历，找右边比自己小的最大者，并交换
交换后当前的index数变大了，要把index+1到末尾的数变成最小值
再从index+1开始，从左向右遍历，进行冒泡排序

时间：O(N^2)
空间：O(1)

优化点：
step1找右边比自己大的最小值时候可以用O(N)时间来实现
step2在做排序的过程，可以证明当前的index+1到末尾一定是降序排列的，所以O(N)时间就可完成排序

具体来说：
1、先找右边比自己大的最小值匹配对(i,j)

		step1: 从n-2开始向左走，若nums[i] < nums[i+1]停止, 此时的i就是要交换的左边
			此时的nums[i+1:n]一定降序，因为不是降序的话就不满足上面的条件
	    step2: 既然nums[i+1:n]为降序，那就从n-1开始向左走，第一个nums[j]比nums[i]大的数就是i右边比它大的最小者
	    step3: swap(i,j)
	        此时的交换后的nums[i+1:n]仍然是降序
	      证明：条件1: nums[j-1] >= nums[j] >= nums[j+1]
	           条件2：             nums[j] > nums[i]

				联立1和2可得：nums[j-1] >= nums[j] > nums[i]
	                ==> nums[j-1] > nums[i]
	          而nums[i]一定大于nums[j+1], 因为step2是满足该条件往左走的
	                ==> nums[i] > nums[j+!]

2、对index+1到n排序
由1可知，index+1到n均为降序，所以双指针交换即可替代排序操作
*/
func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	// step1: 先从n-2开始向左走，找右边比自己大的最小者，找到则结束，记录当前index，找不到index=-1
	index := n - 2
	for index >= 0 {
		t := findLarger(nums, index+1, n-1, nums[index])
		if t != -1 {
			// 交换，并中止
			nums[t], nums[index] = nums[index], nums[t]
			break
		}
		index--
	}
	// step2: 对index+1到n进行排序，从index+1向右走，找右边比自己小的最小者和自己交换，直到末尾结束
	for i := index + 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// 在[l,r]上找严格大于target的最小者，有返回index, 没有返回-1
func findLarger(nums []int, l, r, target int) int {
	index := -1
	curMax := 1<<63 - 1
	for l <= r {
		if target < nums[l] && nums[l] < curMax {
			index, curMax = l, nums[l]
		}
		l++
	}
	return index
}

// 优化之后的法二
func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 || n == 0 {
		return
	}
	// step1: 从右向左找第一个升序索引，例如1,2,4,3 第一个升序索引为1，数字为2,4升序
	index := n - 2
	for index >= 0 && nums[index] >= nums[index+1] {
		index-- // 降序则继续走
	}
	if index >= 0 {
		// index+1到n的降序数组上，从右向左找比index大的最小者
		for j := n - 1; j >= index+1; j-- {
			if nums[j] > nums[index] {
				// 找到交换，并中止
				nums[index], nums[j] = nums[j], nums[index]
				break
			}
		}
	}
	// step2: 对index+1后面进行翻转
	l, r := index+1, n-1
	for l <= r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	nums := []int{1, 3, 2}

	nextPermutation(nums)
	fmt.Println(nums)
}
