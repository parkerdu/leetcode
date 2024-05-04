package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func heOrder(root *TreeNode) {
	if root == nil {
		return
	}
	// step1: 准备一个队列，让root入队
	queue := make([]*TreeNode, 0, 8)
	queue = append(queue, root)
	// step2: 若队列非空，cur出队并打印
	for len(queue) > 0 {
		cur := 	queue[0]
		queue = queue[1:]
		fmt.Println(cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	// step3: cur左孩子先进，右孩子再进
	// step4: 重复step2直到队列为空
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
	//preOrder(root)\
	heOrder(root)
}