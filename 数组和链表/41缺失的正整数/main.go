package main

import "fmt"

func firstMissingPositive(nums []int) int {
	// step1: 遍历数组找最小正整数minP,初始为0
	var minP int
	for i := range nums {
		if nums[i] > 0 && minP == 0 || nums[i] > 0 && nums[i] < minP {
			minP = nums[i]
		}
	}
	// step2: 若minP == 1， 遍历一遍数组，把1放到0index, 把2交换到1index处，然后在从头遍历找到从1递增的最大值，更新为minP
	if minP == 1 {
		i := 0
		// todo 问题：对于nums[1,1]会陷入死循环，无法知道后面的1已经交换了不用再交换了
		// 解决方法：第21行后面的或条件，如果和交换的数相等则继续下一个
		for i < len(nums) {
			if index := nums[i] - 1; 0 <= index && index < len(nums) {
				nums[i], nums[index] = nums[index], nums[i]
				if nums[i] == i+1 || nums[i] == nums[index] { // 解决方法
					i++
				}
			} else {
				i++
			}
		}
		index := 0
		for index < len(nums) && nums[index] == index+1 {
			index++
		}
		minP = nums[index-1]
	} else {
		minP = 0
	}
	// step3: 返回minP+1
	return minP + 1
}

/*
对上面方法做代码上优化：
step1: 先做交换调整位置
step2: 根据调整后的数组来求解，

	若调整后数组的第一个元素就不满足nums[i] = i+1, 则返回i+1就是答案，例如nums[0] = -1 or 6, 返回1

step3: 若遍历完调整后的数组没有返回，说明整个数组都是严格自增1的，例如[1,2,3,4], 此时返回n+1
*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		// 不越界 && 要交互的两个元素不相等 则交互
		// 两个元素相等不交换是应对死循环的
		if index := nums[i] - 1; 0 <= index && index < len(nums) && nums[i] != nums[index] {
			nums[i], nums[index] = nums[index], nums[i]
		} else {
			i++
		}
	}
	for i := range nums {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 2}))
}
