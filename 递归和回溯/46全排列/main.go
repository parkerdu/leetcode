package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	out := make([][]int, 0)
	for i := range nums {
		nums[i], nums[0] = nums[0], nums[i]
		res := permute(nums[1:])
		for j := range res {
			res[j] = append(res[j], nums[0])
			line := make([]int, 1+len(nums))
			copy(line, res[j])
			out = append(out, line)
		}
		nums[i], nums[0] = nums[0], nums[i]
	}
	return out
}

// 空间优化，直接使用nums数组本身交换来做，而不是make和各种append
func permute1(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	dfs(0, n, nums, &res)
	return res
}

func dfs(first, n int, nums []int, res *[][]int) {
	if first == n {
		tmp := []int{}
		for i := range nums {
			tmp = append(tmp, nums[i])
		}
		*res = append(*res, tmp)
		return
	}
	for i := first; i < n; i++ {
		// 当前i和first交换
		nums[i], nums[first] = nums[first], nums[i]
		// 交换之后，first前进继续
		dfs(first+1, n, nums, res)
		// 交换回来
		nums[i], nums[first] = nums[first], nums[i]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
