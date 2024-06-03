package main

// 法一：dfs
func numIslands2(grid [][]byte) int {
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				search(grid,i,j)
				res++
			}
		}
	}
	return res
}

// 从当前i,j节点开始上下左右查找，直到遇到0或者越界结束
func search(grid [][]byte, i,j int) {
	// step1: 边界判断 + 是否为1判断
	if i <0 || i >= len(grid) || j <0 || j>=len(grid[0]) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	// step2: 将当前置为-1
	grid[i][j] = '#'
	// step3: 上下左右查找
	search(grid,i-1,j)
	search(grid,i+1,j)
	search(grid,i,j-1)
	search(grid,i,j+1)
}