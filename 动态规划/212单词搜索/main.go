package main

import (
	"fmt"
)

/*
本题为2024/04/25 pdd一面算法题
我写出了单个单词在board中搜索
 */

func main() {
	board := [][]byte{
		{'o','a','b','n'},
		{'e','t','e','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
		}
	word := "oaba"

	fmt.Println(exist(board,word))
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if find1(board,i,j, word) {
				return true
			}
		}
	}
	return false
}

// 面试时候写出来的代码
func find(board [][]byte, i,j int, word string) bool {
	// step1: 边界条件
	if len(word) == 0 {
		return true
	}
	if i < 0 || i > len(board) {
		return false
	}
	if j < 0 || j > len(board[0]) {
		return false
	}
	// step2: 上下左右递归
	cur := board[i][j] == word[0]
	var up, down, left, right bool
	// 向上寻找
	if i - 1 >= 0 && cur {
		up = find(board, i-1,j, word[1:])
	}
	// 下
	if i + 1 < len(board) && cur {
		down = find(board, i+1,j, word[1:])
	}
	// 左
	if j - 1 >= 0 && cur {
		left= find(board, i,j-1, word[1:])
	}
	// 右
	if j + 1 < len(board[0]) && cur {
		right = find(board, i,j+1, word[1:])
	}
	return up || down || left || right
}


/*
上面版本问题：
1、在字符不相等时候可以直接返回false
2、step2中不用再次判断边界了，step1中都已经处理好了
3、step1的边界条件有问题，i>=len(board)就已经越界了
4、没有对走过的字符做特殊处理，会走出重复路径
 */
func find1(board [][]byte, i,j int, word string) bool {
	// step1: 边界条件
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[0]{
		return false
	}
	// step2: 上下左右递归, 走到这里说明一定是当前字符相等，且在边界内
	// 将当前字符置为一个特殊字符，不是字母
	tmp := board[i][j]
	board[i][j] = 0
	res := find1(board, i-1,j, word[1:]) || find1(board, i+1,j, word[1:]) ||
		find1(board, i,j-1, word[1:]) || find1(board, i,j+1, word[1:])
	board[i][j] = tmp
	return res
}