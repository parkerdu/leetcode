package main

import "fmt"

func mergeSort(arr []int) {
	if arr == nil || len(arr) <2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) {
	if L == R {
		return
	}
	help := make([]int, R-L+1)
	i := 0
	left := L
	right := M + 1
	// 先左右两个指针往右边走
	for left <= M && right <= R {
		if arr[left] <= arr[right] {
			help[i] = arr[left]
			left++
		} else {
			help[i] = arr[right]
			right++
		}
		i++
	}
	// 走完出来，有可能左边还剩，因为右边可能先走完
	for left <= M {
		help[i] = arr[left]
		left++
		i++
	}

	for right <= R {
		help[i] = arr[right]
		right++
		i++
	}
	// 将合并好的数组，复制到原数组中
	for i := range help {
		arr[L+i] = help[i]
	}
}

func main() {
	arr := []int{1, 3, 5, 8, 2, 4, 0}
	mergeSort(arr)
	fmt.Println(arr)
}
