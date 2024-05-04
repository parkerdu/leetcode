package main

import "fmt"

/*
二叉树后序遍历，非递归方法
使用两个辅助栈
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode,0, 8)
	collectStack := make([]int, 0, 8)
	// step1: root节点入栈
	stack = append(stack, root)
	for len(stack) != 0 {
		// step2: 若栈不空，则弹出cur
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// step3: 将cur入collectStack
		collectStack = append(collectStack, cur.Val)
		// step4: cur.left先入栈，再入cur.right
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		// step5: 循环step2，直到栈为空
	}
	// step6: 收集栈中所有元素出栈打印
	for i := len(collectStack) -1; i >= 0; i-- {
		fmt.Print(collectStack[i], ",")
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
	postOrder2023(root)
}

// 手撕后序遍历

func postOrder2023(root *TreeNode) {
	if root == nil {
		return
	}
	// step1: 准备1个辅助栈 stack 和 收集栈 collect
	stack, collect := make([]*TreeNode, 0, 64), make([]*TreeNode, 0, 64)
	// step2: 根节点入stack
	stack = append(stack, root)

	for len(stack) > 0{
		// step3: 若栈非空，弹出一个cur, 并将cur入收集栈
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		collect = append(collect, cur)
		// step4: cur左孩子先进，右孩子再进
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}


	// step5: 打印收集栈中值
	for i := len(collect) - 1; i >= 0; i-- {
		fmt.Println(collect[i].Val)
	}
}