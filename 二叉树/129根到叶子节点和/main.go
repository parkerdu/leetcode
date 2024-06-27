package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：将根到叶子的所有路径作为数组返回上来
func sumNumbers(root *TreeNode) int {
	res := process(root)
	var sum int
	for i := range res {
		for j := range res[i] {
			sum += res[i][j] * int(math.Pow10(j))
		}
	}
	return sum
}

func process(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	left := process(root.Left)
	right := process(root.Right)
	if len(left) == 0 && len(right) == 0 {
		return [][]int{{root.Val}}
	}
	for i := range left {
		left[i] = append(left[i], root.Val)
	}
	for i := range right {
		right[i] = append(right[i], root.Val)
	}
	return append(left, right...)
}

// 法二：对法一空间优化，在递归过程中计算和，每向下一层当前的数*10向下传递
func sumNumbers(root *TreeNode) int {
	return process1(root, 0)
}

func process1(root *TreeNode, add int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return add*10 + root.Val
	}
	return process1(root.Left, add*10+root.Val) + process1(root.Right, add*10+root.Val)
}

func main() {
	one := &TreeNode{Val: 2}
	two := &TreeNode{Val: 3}

	tree := &TreeNode{Val: 1, Left: one, Right: two}
	fmt.Println(sumNumbers(tree))
}
