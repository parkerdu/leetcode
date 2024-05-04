package main

import (
	"fmt"
	"math/rand"
)

func smallSum(arr []int) {
	var sum int
	process(arr, 0, len(arr)-1, &sum)
	fmt.Println(sum)
}

func process(arr []int, L, R int, sum *int){
	if L == R {
		return
	}
	mid := L + ((R -L) >> 1)
	process(arr, L, mid, sum)
	process(arr, mid+1, R, sum)
	merge(arr, L, mid, R, sum)
}

/*
该merge是我第一次自己写的，里面存在一个bug
试着找出来
 */
func merge(arr []int, L, M, R int, sum *int) {
	help := make([]int, R-L+1)
	i, p1, p2 := 0, L, M+1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			*sum += arr[p1] * (R-p2+1)
			help[i] = arr[p1]
			p1 ++
		} else {
			help[i] = arr[p2]
			p2 ++
		}
		i++
	}

	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1 ++
	}

	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2 ++

	}

	for i := range help {
		arr[L+i] = help[i]
	}
}

/*

bug1：左右相等，小和不可相加
bug2: 左右相等，必须先让右指针右移
正确的写法，左侧和右侧比时候，如果左右相等，必须先把右侧拷贝下来，然后右指针p2右移
否则的话若左指针p1右移，会漏掉左边这个数的小和，因为虽然此时左右相等，但是右边p2指针后面还有更大的数比p1指针的数大

例如
左=[1, 2] 右=[1, 4, 5]
第一次p1=p2, 若左指针右移，则会漏掉2个1，因为1<4, 1<5
 */
func merge1(arr []int, L, M, R int, sum *int) {
	help := make([]int, R-L+1)
	i, p1, p2 := 0, L, M+1
	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			*sum += arr[p1] * (R-p2+1)
			help[i] = arr[p1]
			p1 ++
		} else {
			help[i] = arr[p2]
			p2 ++
		}
		i++
	}

	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1 ++
	}
	rand.Int()

	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2 ++

	}

	for i := range help {
		arr[L+i] = help[i]
	}
}

func main() {
	arr := []int{1,1,1,1,13,2,1,5}
	smallSum(arr)

}