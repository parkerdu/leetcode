package main

import (
	"container/heap"
)





// 使用系统自带的heap, 这里实现的是最小堆(就是小根堆)
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mySort(arr []int, k int) {
	h := &IntHeap{}
	copy(*h, arr[0:k])
	heap.Init(h)

	for i := k; i <= len(arr)-1;i++ {
		heap.Push(h, arr[i])
		// 保存了，0,1,2,...,len(arr)-1-k
		arr [i-k] = heap.Pop(h).(int)
	}

	for i:= len(arr)-k; i<len(arr); i++ {
		// 保存，len(arr)-k, ..., len(arr)-1
		arr[i] = heap.Pop(h).(int)
	}
}



func main() {

}