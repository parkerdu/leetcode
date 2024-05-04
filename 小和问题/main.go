package main

import "fmt"

/*
题目描述：小和的定义：给定一个数组[1,2,3,4] 对于每个数将其左边所有比它小的数加起来的和
                             0+1+3+ 6 = 10
*/

func smallSum(nums []int) int {
	/*
		第一次的想法：用mergeSort + 一个数组来做，数组中保存每个元素左边比自己小的值和
		step1: 先求 small数组
		step2: 将small累加
	*/
	n := len(nums)
	small := make([]int, len(nums))
	num := make([][2]int, len(nums))
	for i := range nums {
		num[i] = [2]int{nums[i], i} // 将 数和index重新保存
	}
	process(num, 0, n-1, small)
	var res int
	for i := range small {
		res += small[i]
	}
	return res
}

func process(num [][2]int, l, r int, small []int) {
	if l == r {
		return
	}
	m := (l + r) / 2 // 为防止溢出 写成 l + (r-l)/2
	process(num, l, m, small)
	process(num, m+1, r, small)
	merge(num, l, m, r, small)
}

func merge(num [][2]int, l, m, r int, small []int) {
	p1, p2 := l, m+1
	help := make([][2]int, r-l+1)
	i := 0
	for p1 <= m && p2 <= r {
		if num[p1][0] < num[p2][0] {
			small[num[p2][1]] = num[p1][0] + small[num[p1][1]]
			help[i] = num[p1]
			p1++
		} else {
			help[i] = num[p2]
			p2++
		}
		i++
	}
	for p1 <= m {
		help[i] = num[p1]
		p1++
		i++
	}
	for p2 <= r {
		small[num[p2][1]] = num[m][0] + small[num[m][1]]
		help[i] = num[p2]
		p2++
		i++
	}
	// 拷贝到原数组
	for i := range help {
		num[l+i] = help[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 2}
	fmt.Println(smallSum(nums))
}
