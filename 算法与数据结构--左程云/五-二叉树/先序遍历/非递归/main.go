package main

import "fmt"

/*
二叉树先序遍历，非递归方法
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode,0, 8)
	// step1: root节点入栈
	stack = append(stack, root)
	for len(stack) != 0 {
		// step2: 若栈不空，则弹出cur
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// step3: 打印cur.val
		fmt.Print(cur.Val, ",")
		// step4: cur.right先入栈，再入cur.left
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
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
	preOrder2023(root)
}



/*
2023.11.22 二刷
 */

func preOrder2023(root *TreeNode) {
	if root == nil {
		return
	}
	// step1: 准备1个栈
	stack := make([]*TreeNode, 0, 64)
	// step2: 将根节点进栈
	stack = append(stack, root)
	// step3: 若栈非空则循环遍历
	for len(stack) > 0 {
		// step3.1: 弹出一个元素并打印
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(cur.Val)
		// step3.2: 右孩子先进，左孩子再进
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
}