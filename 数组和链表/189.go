package main

// 21:48 - 22:18 ac 法二 5%,21%

/*
法1： 开辟一个数组，拼接
法2： 从n-k-1开始，先拿出这个数，然后前面整体向后移动1位，再将当前数插入最前面，时间(n-k)*k, 空间O(1)
 */
func rotate(nums []int, k int)  {
	n := len(nums)
	if k >= n {
		k = k % n
	}
	if k == 0 {
		return
	}
	for i := n-k; i <= n-1; i++{
		cur := nums[i]
		for j := i; j > i-n+k; j -- {
			nums[j] = nums[j-1]
		}
		nums[i-n+k] = cur
	}
}

// 看了答案 知道法三
/*
三次反转 用时O(n)
 */
func rotate(nums []int, k int)  {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	reverse(nums,0,n-1)
	reverse(nums,0,k-1)
	reverse(nums,k,n-1)
}

func reverse(nums []int, l,r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}