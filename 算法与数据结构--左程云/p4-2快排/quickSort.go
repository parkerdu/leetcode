package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	quick(arr, 0, len(arr)-1)

}

func quick(arr []int, L, R int) {
	if L < R {
		// 生产[L, R]上的随机index
		randomIndex := L + int(rand.Float32() *float32(R -L +1))
		// 将随机数和最后一个数交换
		swap(arr, randomIndex, R)
		left, right := partion(arr, L, R)
		quick(arr, L, left-1)
		quick(arr, right+1, R)
	}
}

/*
荷兰国旗问题，取num=arr[R], 将arr分为三个区
1、<num区 index[L, left-1]
2、=num区 index[left, right]
3、>num区 index[right+1, R]
返回中间=num区的，左右index
 */

// 第一次写出，有bug

// 思考：为什么要先对[L, R-1]交换，最后再交换一个R, 不能直接对[L, R]进行吗
// 我的理解，老师写的没有开辟num这个变量，直接拿arr[R]比较的，所以要最后交换
// 而我的写法新开辟一个num保存了arr[R], 所以理论上是可以直接对[L,R]进行的
func partion(arr []int, L, R int) (left, right int){
	left, right = L, R-1
	num := arr[R]
	i := L
	// 该for-loop只完成了，闭区间[L, R-1]上的交换，因此最后要把R上的数交换
	for i < R {
		if arr[i] < num {
			swap(arr, i, left)
			i++
			left ++
		} else if arr[left] == num {
			i++
		} else {
			swap(arr, i,  right)
			right--
			i++
		}
	}
	swap(arr, R, right+1)
	return left, right + 1
}

// 上面的partion有bug, 就是和右边交换的时候，i不能动，因为右边交换过来的数还没有比较
func partion1(arr []int, L, R int) (left, right int){
	left, right = L, R-1
	num := arr[R]
	i := L
	for i < R {
		if arr[i] < num {
			swap(arr, i, left)
			i++
			left ++
		} else if arr[left] == num {
			i++
		} else {
			swap(arr, i,  right)
			right--
		}
	}
	swap(arr, R, right+1)
	return left, right + 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
 }

func main() {
	arr := []int{1, 4, 3, 2, 5, 5, 78, 0, 3, 4}
	quickSort(arr)
	fmt.Println(arr)
}