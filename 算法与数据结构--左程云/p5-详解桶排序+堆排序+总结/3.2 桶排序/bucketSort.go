package main

import (
	"fmt"
	"math"
	"sort"
)

func bucketSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	bigBit := maxBits(arr)
	bucketSort1(arr, 0, len(arr)-1, bigBit)
}

// 对arr[l, r]上进行桶排序
func bucketSort1(arr []int, l, r, bigBit int) {
	radix := 10 // 基数为10， 桶排序也叫基数排序
	help := make([]int, r-l+1)
	for d := 1; d <= bigBit; d++ {
		count := make([]int, radix)
		// 假设d=1, 则意思为，统计arr中个位数的词频，并放入count数组中
		for i :=l; i <= r; i++ {
			bitNum := getBitNum(arr[i], d)
			count[bitNum]++
		}

		// 假设d=1， 将count数组意思改写前缀和，count[i] = 数组arr中个位数<=i个数
		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]
		}

		// 从右向左遍历arr, 将arr中元素按个位上数排序（相当于进一次桶+出一次桶，后进后出）
		for j := r; j >=l; j--{
			bitNum := getBitNum(arr[j], d)  // 获取arr[j]的个位数，假设arr[j]=153, 则bigNum=3
			index := count[bitNum] - 1      // 看下<=arr[j]个位数的有几个, 假设bigNum=3, count[3]=4, 则意思为小于等于3的有4个
			help[index] = arr[j]   // 将 arr[j] 放到 help数组中的第3位
			count[bitNum]--
		}

		// 再把help上的数拷贝到原数组arr中
		for i := range help {
			arr[l+i] = help[i]
		}
	}
}


// 取出num这个数上的第几位（从低到高）
// 例如getBitNum(15, 1) 表示去除15上的第一位（个数），return 5
/*
因此
要取第一位（个位）， num右移0位，再取模
要取第二位（十位）， num右移1位，再取模
要取第三位（百位）， num右移2位，再取模
 */

func getBitNum(num, bit int ) int{
	return (num / int(math.Pow10(bit-1))) % 10
}


// 一个数组arr, 最大数有多少位
func maxBits(arr []int) int {
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	bigBit := 0
	for max > 0 {
		max = max / 10
		bigBit++
	}
	return bigBit
}

func main() {
	sort.Sort()
	arr := []int{1, 3, 15, 17, 29, 30, 120,121, 19}
	bucketSort(arr)
	fmt.Println(arr)
}
