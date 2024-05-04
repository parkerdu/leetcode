package main

import "fmt"

/*
没写出来，看评论后自己写的
思路：
nums = [1,      2,    3,    4]
left   [?,      1,   1*2, 1*2*3] 从左向右便利
right  [4*3*2  4*3    4     ?]
?处补1,再将left和right相乘即为答案
*/

// 时间O(N), 空间O(N)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{0}
	}
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 2; i >= 0; i-- {
		right = nums[i+1] * right
		res[i] = res[i] * right
	}
	return res
}

func main() {
	nums := []int{1, 0, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
