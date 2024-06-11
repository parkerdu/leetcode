package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	process(matrix, 0, 0, 0, &res)
	return res
}

// direct: 本次走完，下一个往哪里走
func process(matrix [][]int, i, j int, direct int, res *[]int) int {
	// res长度够了，退出
	if len(*res) == len(matrix)*len(matrix[0]) {
		return direct
	}
	// base, 越界，或回头路，转变方向返回
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || matrix[i][j] == 101 {
		return changeDirect(direct)
	}
	*res = append(*res, matrix[i][j])
	matrix[i][j] = 101
	// 根据当前方向，和下一个元素越界，回头路情况，决定下一次的方向
	var newDirect int
	x, y := path(direct, i, j)
	newDirect = process(matrix, x, y, direct, res)
	// 发生方向变更，本次路重走
	if direct != newDirect {
		// 根据新方向决定走哪里
		x, y := path(newDirect, i, j)
		process(matrix, x, y, newDirect, res)
		return newDirect
	}
	return direct
}

func path(direct, i, j int) (int, int) {
	if direct == 0 {
		return i, j + 1
	}
	if direct == 1 {
		return i + 1, j
	}
	if direct == 2 {
		return i, j - 1
	}
	if direct == 3 {
		return i - 1, j
	}
	return i, j
}

// 0右，1下，2左，3上
func changeDirect(direct int) int {
	return (direct + 1) % 4
}

func main() {
	m := make([][]int, 3)
	m[0] = []int{1, 2, 3, 4}
	m[1] = []int{5, 6, 7, 8}
	m[2] = []int{9, 10, 11, 12}
	fmt.Println(spiralOrder1(m))
}

/*
法二：自己想的4指针+1方向
4指针：up,down,left,right 代表矩阵中的索引
当direct为向右走时候，在up行上面，从left指针开始遍历到right指针结束，该行遍历玩，up++
当direct为向下走时候，在right列上面，从up开始到down结束，走完向下的一列后，right--
direct 向左，在down行上利用，left,right信息向左走

direct利用取模来控制方向
循环结束条件：res数组长度 == m*n
*/
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	res := make([]int, 0)
	var direct int
	for len(res) != len(matrix)*len(matrix[0]) {
		// 向右走
		if direct == 0 {
			res = append(res, matrix[up][left:right+1]...)
			up++
		}

		// 向下
		if direct == 1 {
			for i := up; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		// 向左
		if direct == 2 {
			for j := right; j >= left; j-- {
				res = append(res, matrix[down][j])
			}
			down--
		}
		if direct == 3 {
			// 向上
			for i := down; i >= up; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
		direct = (direct + 1) % 4
	}
	return res
}
