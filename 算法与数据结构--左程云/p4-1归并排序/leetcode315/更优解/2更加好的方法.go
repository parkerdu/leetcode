package main

import "fmt"

/*
其实这个题跟最小和问题不一样的地方就是要知道排序后数组的元素对应原数组中的index是什么
那么我们可以稍微改造一下要排序的内容，既然想知道某个元素对应在原数组中的index，那我就构造一个数组a=[原数组中元素，该元素的index]
例如nums = [5,2,6,1] 修改为 arr = [[5,0], [2,1], [6,2], [1,3]]
对arr进行归并排序，这样就解决了merge过程中无法知道某个元素对应原数组index问题
 */

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))
	if nums == nil || len(nums) < 2 {
		return count
	}

	// 修改原数组，为带index的数组
	arr := make([][2]int, len(nums))
	for i, v := range nums {
		arr[i] = [2]int{v, i}
	}
	process(arr, count, 0, len(nums)-1)
	return count
}

func process(arr [][2]int, count []int, L, R int) {
	if L == R {
		return
	}
	mid := L +((R-L) >>1)
	process(arr, count, L, mid)
	process(arr, count, mid + 1, R)
	merge(arr, count, L, mid, R)
}

func merge(arr [][2]int,  count []int, L, M, R int) {
	help := make([][2]int, R-L+1)
	p1, p2 := L, M+1
	i := 0
	for p1 <= M && p2 <= R{
		if arr[p1][0] > arr[p2][0] {
			count[arr[p1][1]] += R-p2+1
			help[i] = arr[p1]
			i++
			p1++
		} else {
			help[i] = arr[p2]
			i++
			p2 ++
		}
	}

	for p1 <= M {
		help[i] = arr[p1]
		p1++
		i++
	}


	for p2 <= R{
		help[i] = arr[p2]
		p2++
		i++
	}

	for i := range help {
		arr[L+i] = help[i]
	}
}

func main() {
	arr := []int{5,2,6,1}
	fmt.Printf("结果为count：%v", countSmaller(arr))

}