package main

import "fmt"

func longestConsecutive(nums []int) int {
	/*
		第一次没想出来，看了答案后的思路：
		将nums构造成一个哈希表，利用哈希表的O(1)查询特性，先找连续数组的起点，再看能坚持多长
		时间复杂度O(N) ：举个例子 哈希表中有3,2,1， 遍历
		对于3，发现3-1=2存在，所以3不是起点，什么都不做继续
		对于2：发现1存在，继续
		对于1，发现0不存在，说明1是起点，初始化当前len=1,继续看2也在len=2, 再看3也在len=3
	*/
	if len(nums) == 0 {
		return 0
	}
	var res int
	// step1: 构造哈希表
	set := make(map[int]struct{})
	for i := range nums {
		// 在set中跳过，不在创建key
		_, ok := set[nums[i]]
		if !ok {
			set[nums[i]] = struct{}{}
		}
	}
	// step2: 遍历哈希表，先找起点，找到起点再找能持续多长
	for k := range set {
		_, ok := set[k-1]
		// k-1不在里面表示 起点开始寻找长度
		if !ok {
			len := 1
			for {
				// k+1是否存在，存在则len+1, 不存在比较并返回
				if _, ok := set[k+1]; ok {
					len++
					k++
				} else {
					if len > res {
						res = len
					}
					break
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{100, 1, 4, 200, 2, 3}
	fmt.Println(longestConsecutive(nums))

}
