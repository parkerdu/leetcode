package main

import "fmt"

/*
题目5
节点数为n高度不大于m的二叉树个数
现在有n个节点，计算出有多少个不同结构的二叉树
满足节点个数为n且树的高度不超过m的方案
因为答案很大，所以答案需要模上1000000007后输出
测试链接 : https://www.nowcoder.com/practice/aaefe5896cce4204b276e213e725f3ea

 */


// 法1, 递归
func maxNumTree(n,m int) int{
	mod := 1<<9+7
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, m+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return f1(cache, n,m, mod)
}

func f1(cache [][]int, n,m, mod int) int {
	// step1: 边界条件
	if n == 0 { // 节点数为0,1中
		return 1
	}
	if m == 0 {  // 节点数不为0，高度为0,0种
		return 0
	}
	// step2: 缓存有，从缓存拿
	if cache[n][m] != -1 {
		return cache[n][m]
	}
	// step3: 计算0,1，2，...k (k=n-1)的最大二叉树数量
	// f(n,m) = E(k=0到n-1) f1(k, m-1) * f1(n-1-k,m-1)
	var ans int
	for k := 0; k < n; k++{
		ans = (ans +  (f1(cache, k, m-1, mod) * f1(cache, n-1-k, m-1, mod)) % mod) % mod
	}
	cache[n][m] = ans
	return ans
}


// 法二： 动态规划
func maxNumTree2(n,m int) int{
	// step1： 开辟1个数组，长度为n+1代表某一列的数据
	dp := make([]int, n+1)
	// step2: 边界条件dp[0] = 1, 第一列除了首个元素外全为0
	dp[0] = 1
	// step3: j=1从第1列开始，每列安装从下向上更新i=n,n-1,...1
	for j :=1; j <= m; j++ {
		for i :=n ;i >=1; i-- {
			dp[i] = 0  // 先将原来的dp[i]重置，反正我也不依赖它
			for k :=0; k<i;k++ {
				dp[i] = dp[i] + dp[k] * dp[i-1-k]
			}
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(maxNumTree2(3,3))
}