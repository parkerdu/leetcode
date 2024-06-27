package main

/*
本题dp的关键是dp[i,j]的定义
1、定义dp[i][j]为matrix[i][j]为右下角的正方形最大边长
2、递推公式：

	i-1,j-1     i-1,j
	  i,j-1       i,j

对于上面的i,j的最大边长，和左上角，左边，上面三者息息相关
先看i-1,j如果他的边长是1，i,j想和他组合起来成为一个正方形，那么剩下两个的边长也至少是1

	如果他的边长是2，i,j想和他组合起来成为一个正方形边长为3，那么两位两个也至少是2

如果3个中任意一个为0，答案为1
如果3个中最小值为1，那么可以看成3个1，答案就是2
所以可以推出：dp[i][j] = 1 + min{左上角，左边，上面}
3、初始条件：第一行和第一列可以先计算
4、方向，每行从左到右，从第1行开始
*/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var edge int
	dp := make([]int, len(matrix[0]))
	for i := range matrix[0] {
		if matrix[0][i] == '1' {
			dp[i] = 1
			edge = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		leftUp := dp[0]
		dp[0] = int(matrix[i][0] - '0')
		if dp[0] > edge {
			edge = dp[0]
		}
		for j := 1; j < len(matrix[i]); j++ {
			tmp := dp[j]
			if matrix[i][j] == '1' {
				dp[j] = min(leftUp, dp[j-1], dp[j]) + 1
				if dp[j] > edge {
					edge = dp[j]
				}
			} else {
				dp[j] = 0
			}
			leftUp = tmp
		}
	}
	return edge * edge
}
