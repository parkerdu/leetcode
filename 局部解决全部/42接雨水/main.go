package main

import "fmt"

/*
13分钟后的思路：
分层来看：每一层只看一个高度上能存多少水，第一层计算完之后，整体高度下降1，再看能存多少水，重复最大高度后停止
每一次只找0，然后从0开始左右向外扩展，直到第一个非0的数后停止，得到right-left为本轮可以接到的雨水
step1: 找出数组的最大值max
step2: 循环max次
	step2.1 从左到右遍历height, 发现有0的数，则开始左右外扩，找到
	step2.2 整体高度下降1
 */
func trap(height []int) int {
	// step1: 找出最大值
	max := 0
	for i := range height {
		if height[i] > max {
			max = height[i]
		}
	}
	var res int
	// step2: 做max次循环
	for i :=0; i<max; i++ {
		j := 0
		for {
			if j >= len(height) {
				break
			}
			if height[j] == 0 {
				product, nextIndex := findAndCompute(height,j)
				if product > 0 {
					res += product
					j = nextIndex
					continue
				}
			}
			j++
		}
		// 所有高度下降1
		for j := range height {
			if height[j] >= 1 {
				height[j]--
			}
		}
	}
	return res
}

// 从当前index开始，找到左右第一个非0的索引，找到返回r-l, 找不到返回0
func findAndCompute(height []int, index int) (int,int) {
	left, right := index,index
	for left >=0 && right < len(height) && (height[left] == 0 || height[right] == 0) {
		if height[left] != 0 && height[right] != 0 {
			break
		}
		if height[left] == 0 {
			left--
		}
		if height[right] == 0 {
			right++
		}
	}
	// 越界则返回0，不越界说明找到了
	if left <0 || right >= len(height) {
		return 0,0
	}
	return right-left-1, right
}



// 看了解答后代码：https://zhuanlan.zhihu.com/p/107792266
// 法1： 时间O(N2), 空间O(1
func trap1(height []int) int {
	/*
	法一：时间O(N**2), 每次遍历时候找左右最大值
	 */
	var res int
	for i := 1; i< len(height);i++{
		l,r := findMaxHeight(height,i)
		min := l
		if min > r {
			min = r
		}
		res += min - height[i]
	}
	return res
}

// 找leftMax为从0开始到index的最大值，包含index = height[0:index+1] 的最大值
// rightMax 为从index开始一直到后面的最大值，包含index
func findMaxHeight(height []int, index int) (int,int) {
	var left,right int
	for i := 0; i <= index; i ++ {
		if height[i] > left {
			left = height[i]
		}
	}
	for i := index; i < len(height); i++ {
		if height[i] > right {
			right = height[i]
		}
	}
	return left,right
}

// 法二：缓存优化 时间0(N), 空间O(N), 缓存优化
func trap2(height []int) int {
	/*
		法二：时间O(N), 每次遍历时候找左右最大值
		定义2个数组：leftMax 和 rightMax,
	   其中leftMax[i]代表h[i] 左边从0到i包含i的最大值
	  其中rightMax[i]代表h[i] 右边从i到最后包含i的最大值
	  先把这两个数组计算好，再用
	*/

	// step1: 计算左右数组
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0], rightMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = leftMax[i-1]
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		}
	}
	for j := n-2; j >=0; j-- {
		rightMax[j] = rightMax[j+1]
		if height[j] > rightMax[j+1] {
			rightMax[j] = height[j]
		}
	}


	var res int
	for i := 1; i< len(height);i++{
		l,r := leftMax[i], rightMax[i]
		min := l
		if min > r {
			min = r
		}
		res += min - height[i]
	}
	return res
}

// 法三：时间0(N), 空间O(1) 双指针
func trap3(height []int) int {
	/*
	法二：时间O(N),
	两个指针，left和right, left从1开始递增，right从n-2开始递减
	定义 curLeftMax 为[0,left]闭区间最大值， curRightMax 为[right,-]闭区间最大值

	左右两个最大值，哪个小就计算哪边
	*/
	var res int
	n := len(height)
	if n <= 2 {
		return 0
	}
	left, right := 1, n-2
	curLeftMax, curRightMax := height[0], height[n-1]
	for left <= right {
		// 先计算当前的左右最大值
		if height[left] > curLeftMax {
			curLeftMax = height[left]
		}
		if height[right] > curRightMax {
			curRightMax = height[right]
		}
		// 找左右最小者，计算当前的能接的雨水
		if curLeftMax < curRightMax {
			res += curLeftMax - height[left]
			left++
		} else  {
			res += curRightMax - height[right]
			right--
		}
	}
	return res
}

func main() {
	h := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap2(h))
}