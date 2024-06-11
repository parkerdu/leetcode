package main

// 法一：dfs
func numIslands2(grid [][]byte) int {
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				search(grid, i, j)
				res++
			}
		}
	}
	return res
}

// 从当前i,j节点开始上下左右查找，直到遇到0或者越界结束
func search(grid [][]byte, i, j int) {
	// step1: 边界判断 + 是否为1判断
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	// step2: 将当前置为-1
	grid[i][j] = '#'
	// step3: 上下左右查找
	search(grid, i-1, j)
	search(grid, i+1, j)
	search(grid, i, j-1)
	search(grid, i, j+1)
}

// 法二：并查集, 对每个1建立独立的集合，记录当前的连通分支数量，然后遍历数组，右边和下面是1则合并（上和左是重复动作）
// 合并完成后当前的连通分支数量即为答案
func numIslands3(grid [][]byte) int {
	u := newUnionFind(grid)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '1' {
				continue
			}
			if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
				u.unionSet(u.index(i, j), u.index(i, j+1))
			}
			if i+1 < len(grid) && grid[i+1][j] == '1' {
				u.unionSet(u.index(i, j), u.index(i+1, j))
			}
		}
	}
	return u.Branches
}

type union struct {
	Father   []int
	Branches int //连通分支数
	column   int // 一列多少个元素
}

// 从cur开始一直向上找父亲，并做扁平化
func (u *union) find(cur int) int {
	if cur != u.Father[cur] {
		u.Father[cur] = u.find(u.Father[cur])
	}
	return u.Father[cur]
}

func (u *union) isSame(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *union) unionSet(x, y int) {
	fx, fy := u.find(x), u.find(y)
	if fx != fy {
		u.Father[fy] = fx
		u.Branches--
	}
}

func (u *union) index(i, j int) int {
	return i*u.column + j
}

func newUnionFind(grid [][]byte) *union {
	m, n := len(grid), len(grid[0])
	u := &union{
		Father:   make([]int, m*n),
		Branches: 0,
		column:   n,
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				u.Father[u.index(i, j)] = u.index(i, j)
				u.Branches++
			}
		}
	}
	return u
}
