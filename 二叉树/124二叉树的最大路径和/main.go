package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	inMax, outMax, _ := process(root)
	return max(inMax, outMax)
}

// 从当前root开始的最大值，返回包不包含root
// 返回值：含根最大值，不含根最大值, max(根+左，根+右，根)
func process(root *TreeNode) (int, int, int) {
	if root == nil {
		return -1001, -1001, -1001
	}
	includL, left, pl := process(root.Left)
	includR, right, pr := process(root.Right)
	// root不参与
	outcludMax := max(left, right, includL, includR)
	// root只参与左或者右， 返回这个信息是为了让自己的父亲能利用该信息
	path := max(root.Val+pl, root.Val+pr, root.Val)
	// 当root参与其中，root一旦参与左右孩子都必须是含根的单条路，否则就违背了题目path的特性
	// 一个path的要求是一个节点只能有两条边，出现下面的3条边就不是path
	/*
	   1
	     \
	      3   // 这里的节点3存在3个边所以不是path, 在向节点1汇报时候，只能返回3+5=8的path
	     /   \
	   4      5
	*/
	includMax := max(path, root.Val+pl+pr)
	return includMax, outcludMax, path
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	left := &TreeNode{Val: -5}
	right := &TreeNode{Val: 7, Left: one, Right: two}
	tree := &TreeNode{Val: -1, Left: left, Right: right}

	fmt.Println(maxPathSum(tree))
}
