package main

import "fmt"

func heapSort(list []int) {
	// step1: 构造大根堆
	heapSize := 0
	for i := range list {
		heapInsert(list, i)
		heapSize++
	}
	// step2: 每次弹出堆顶元素，和最后一个元素交换，heapsize--
	for heapSize > 0 {
		swap(list, 0, heapSize-1)
		heapSize--
		heapify(list, 0, heapSize)  // 重新从顶向下堆化
	}
}


func heapify(list []int, index int, heapSize int) {
	// step1: 左是否存在，存在则max = left
	for index < heapSize {
		left := 2 *index+1
		if left >= heapSize {
			break
		}
		max := left
		if left + 1< heapSize && list[left+1] > list[left] {
			max = left +1
		}
		if list[max] > list[index] {
			swap(list, max, index)
			index = max
		} else {
			return
		}
	}
	// step2: 右是否存在，存在和max比较
	// step3: 当前index 和左右最大值比较，比index大则交换，index = max重复
}
// 往当前堆中插入一个元素，向上比较，直到为root或者比父小
func heapInsert(list []int, index int) {
	for list[index] > list[(index - 1) / 2 ] {
		swap(list, index, (index - 1) / 2)
		index = (index - 1) / 2
	}
}

func swap(list []int, l,r int) {
	list[l], list[r] = list[r], list[l]
}

func main() {
	l := []int{2,1,5,7,3}
	heapSort(l)
	fmt.Println(l)
}