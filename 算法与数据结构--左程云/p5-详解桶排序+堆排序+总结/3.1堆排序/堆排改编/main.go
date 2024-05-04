package main

import (
	"fmt"
)

/*
题目描述：
给定一个数组，将其按照从小到大顺序排序，打印出该数组排序后的下标
例如arr=[6, 1, 4, 3, 2 ], 排序后为[1, 2, 3, 4, 6] 输出结果：[1, 4, 3, 2, 0]
 */


func heapSort(arr [][2]int) {
	if len(arr) < 2 {
		return
	}

	// step1: 将所有数组元素变成一个大根堆
	for i := range arr {
		heapInsert(arr, i)
	}

	// step2: 将最大数和堆得最后一个数交换，并和堆最大数断开链接，使用heapify，直到全部交换完成
	heapSize := len(arr)
	for heapSize > 0 {
		swap(arr, 0 ,heapSize-1)
		heapSize--
		heapify(arr, 0, heapSize)
	}
}

func heapInsert(arr [][2]int, index int) {
	// 定义：对于一个大根堆，从其尾部插入index的数使其仍然为大根堆
	// 实现：对于新来的index, 让其一直往上找其父亲（index-1）/2 pk, 知道打不过父亲了或者已经杀到根节点了
	for arr[(index-1)/2][0] < arr[index][0] {
		swap(arr, index, (index-1)/2)
		index = (index-1)/2
	}
}



func heapify(arr [][2]int, index, heapSize int) {
	// 定义：对于一个大根堆arr, 将其index位置上的元素更换掉了， 使其从index往下开始仍然是一个大根堆
	// 实现：从index开始和左右子树的最大者比较，干不过儿子的一直往下走，直到没有儿子干得过他或者走到叶子节点了
	// 时间复杂度： O(logN)
	left := 2*index +1
	for left < heapSize {
		max := left
		if (left + 1) < heapSize && arr[left+1][0] > arr[left][0] {
			max = left+1
		}
		if arr[index][0] < arr[max][0] {
			swap(arr, max, index)
			index = max
			left = 2 * index + 1
		} else {
			return
		}
	}
}


func swap(arr [][2]int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}


func solution(arr []int) {
	// step1: 将arr 变成带下标的数组
	indexArray := make([][2]int, 0, len(arr))
	for index, value := range arr {
		indexArray = append(indexArray, [2]int{value, index})
	}

	// step2: 将带下标的数组进行堆排序
	heapSort(indexArray)

	// step3: 获取排序后元素下标
	for i := range indexArray{
		fmt.Println(indexArray[i][1])
	}
}

func main() {
	arr := []int{6, 1, 4, 3, 2}
	solution(arr)
}