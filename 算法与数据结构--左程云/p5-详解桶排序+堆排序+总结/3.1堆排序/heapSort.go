package main

import "fmt"

func heapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 步骤一：将数组建为大根堆，时间复杂度O(N*logN)
	for i := range arr {
		heapInsert(arr,i)
	}

	// 步骤一，让整个数组变成堆得优化方法，时间复杂度O(N)
	//n := len(arr) -1
	//heapSize := len(arr)
	//for n >= 0 {
	//	heapify(arr, n, heapSize)
	//	n--
	//}

	heapSize := len(arr)
	// 步骤二：每次交换arr[0]和arr[heapSize-1], 将剩下数堆化
	for heapSize > 0 {
		swap(arr, 0, heapSize-1)
		heapSize--
		heapify(arr, 0, heapSize)
	}
}


// 时间复杂度O(logN)
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index-1)/2
	}
}

func heapify(arr []int, index, heapSize int) {
	left := 2*index +1
	for left < heapSize {
		max := left
		if (left +1) < heapSize && arr[left+1] > arr[left] {
			max = left+1
		}
		if arr[max] > arr[index] {
			swap(arr, max, index)
			index = max
			left = 2*index+1
		} else {
			return
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{1, 4,12,4,5,34,67,0}
	heapSort(arr)
	fmt.Println(arr)
}