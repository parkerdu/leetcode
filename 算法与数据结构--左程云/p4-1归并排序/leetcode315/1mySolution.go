package main

import "fmt"


/*
我的解法：
利用逆序mergeSort（从大到小）,
一、但开辟一个index数组用来保存排序后数组对应的原nums中的index
例如
原数组nums:[2,5,6,1]
排序后数组：[6,5,2,1]
则index： [2,1,0,3] 2表示排序后数组中的6，对应nums中的第二个元素
二、开辟一个map[int]int {index: 逆序对个数}
	key为nums下标，value为该nums[key]的逆序对数

在merge过程中改变map[key]
1、当arr[p1] > arr[p2]时，map[index[p1]] += R-p2+1
	并且让大的数，arr[p1]先移到help数组中，然后p1右移
2、当arr[p1] <= arr[p2]时，没有，p2右移
最后当更新原数组arr时候，同时更新index数组
 */
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	if nums == nil || len(nums) < 2 {
		return res
	}
	count := make(map[int]int)  // 因为最终要输出的逆序对顺序是按原数组来输出的，所以开辟一个map{"原数组下标”："该下标元素对应的逆序对数"}
	index := make([]int, len(nums))  // 开辟一个index用来记录，排序后数组对应原数组的下标
	for i := range nums{
		index[i] = i
	}
	process(nums, index, count, 0, len(nums)-1)
	fmt.Printf("排序后数组为：%v", nums)
	fmt.Printf("排序后index为：%v", index)
	for k, v := range count{
		res[k] = v
	}
	return res
}

func process(arr, index []int, count map[int]int, L, R int) {
	if L == R {
		return
	}
	mid := L + ((R-L) >> 1)
	process(arr, index, count, L, mid)
	process(arr, index, count, mid+1, R)
	merge(arr,index, count, L, mid, R)
}

func merge(arr, index []int,count map[int]int, L, M, R int) {
	helpArr := make([]int, R-L+1)
	helpIndex := make([]int, R-L+1)
	p1, p2 := L, M+1
	i := 0
	for p1 <= M && p2 <= R {
		if arr[p1] > arr[p2] {
			// 右边比左边小
			count[index[p1]] +=  R-p2+1  // 产生逆序对，则该元素对应的下标对应value加上右边的比它小的个数
			helpArr[i] = arr[p1]
			helpIndex[i] = index[p1]
			p1++
			i++
		} else {
			helpArr[i] = arr[p2]
			helpIndex[i] = index[p2]
			i++
			p2++
		}
	}

	for p1 <= M {
		helpArr[i] = arr[p1]
		helpIndex[i] = index[p1]
		i++
		p1 ++
	}

	for p2 <= R {
		helpArr[i] = arr[p2]
		helpIndex[i] = index[p2]
		i++
		p2++
	}

	for i := range helpArr {
		arr[L+i] = helpArr[i]
		index[L+i] = helpIndex[i]
	}
}

func main() {
	arr := []int{5,2,6,1}
	fmt.Printf("结果为count：%v", countSmaller(arr))

}