package main

import "fmt"

/*
二叉树中序遍历，非递归方法
对每个子树都有：左—根-右
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode,0, 8)
	// step1: root左边界入栈
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		// step2: 若栈不空，则弹出cur
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// step3: 打印cur.val
		fmt.Print(cur.Val, ",")
		// step4: 对cur.right这颗子树，左边界全部入栈
		right := cur.Right
		for right != nil {
			stack = append(stack, right)
			right = right.Left
		}
		// step5: 循环step2，直到栈为空
	}
}

func main() {
	node6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		Val:   4,
		Left:  node6,
		Right: node5,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  node3,
		Right: nil,
	}

	root := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node4,
	}
	/*
		二叉树为
				 1
		     2        4
		  3	       6      5
	*/
	midOrder(root)
}