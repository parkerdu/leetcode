package main

import (
	"fmt"
)

// 开始21:01
func numIslands(grid [][]byte) int {
	// 开一个数组arr, arr[i] 代表第i个联通分支，返回len(arr)
	arr := make(map[int]map[[2]int]struct{})
	key := 1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			// step1:. 为1判断当前index的 上和左 在不在arr的任一集合中，在则加入，不在step2
			// 这里有问题：左上有1则直接和他合并就行了，不需要去查是否存在的
			var left, up int
			if i-1 >= 0 {
				left = isExist(arr, [2]int{i - 1, j})
			}
			if j-1 >= 0 {
				up = isExist(arr, [2]int{i, j - 1})
			}
			if left != 0 && up != 0 && left != up {
				// 合并两个集合，这里需要做大量内存迁移操作，还需要删除arr1个元素
				// 合并到left, 将up置空
				for k := range arr[up] {
					arr[left][k] = struct{}{}
				}
				delete(arr, up)
			}
			if left != 0 {
				arr[left][[2]int{i, j}] = struct{}{}
				continue
			}
			if up != 0 {
				arr[up][[2]int{i, j}] = struct{}{}
				continue
			}

			// step2: 左和上都没有1， 那么创建一个新集合，保存当前位置
			set := make(map[[2]int]struct{})
			set[[2]int{i, j}] = struct{}{}
			arr[key] = set
			key++
		}
	}
	return len(arr)
}

// key 在 arr中的哪个map中，返回其key
func isExist(arr map[int]map[[2]int]struct{}, key [2]int) int {
	for i := range arr {
		if _, ok := arr[i][key]; ok {
			return i
		}
	}
	return 0
}

func main() {
	grid := [][]uint8{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
	}
	fmt.Println(numIslands1(grid))
}

//

func numIslands1(grid [][]byte) int {
	u := newUnionFind(grid)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			if j-1 >= 0 && grid[i][j-1] == '1' {
				u.union(u.index(i, j), u.index(i, j-1))
			}
			if i-1 >= 0 && grid[i-1][j] == '1' {
				u.union(u.index(i, j), u.index(i-1, j))
			}
		}
	}
	return u.branches
}

type unionFind struct {
	father   []int
	size     []int
	branches int // 连通分支数
	columns  int // 一列多少个元素
}

func (u *unionFind) index(i, j int) int {
	return i*u.columns + j
}

// 二维数组使用技巧一维化，对每个1建立一个集合
func newUnionFind(arr [][]byte) *unionFind {
	u := &unionFind{}
	m, n := len(arr), len(arr[0])
	u.father = make([]int, m*n)
	u.size = make([]int, m*n)
	u.columns = n
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == '1' {
				u.father[u.index(i, j)] = u.index(i, j)
				u.size[u.index(i, j)] = 1
				u.branches++
			}
		}
	}
	return u
}

func (u *unionFind) find(cur int) int {
	help := make([]int, 0, 48)
	for cur != u.father[cur] {
		cur = u.father[cur]
		help = append(help, cur)
	}
	// cur为当前代表，将路径中元素扁平化
	for i := range help {
		u.father[help[i]] = cur
	}
	return cur
}
func (u *unionFind) isSame(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *unionFind) union(a, b int) {
	ha, hb := u.find(a), u.find(b)
	if ha == hb {
		return
	}
	if u.size[ha] < u.size[hb] {
		u.father[ha] = hb
		u.size[hb] += u.size[ha]
	} else {
		u.father[hb] = ha
		u.size[ha] += u.size[hb]
	}
	u.branches--
}
