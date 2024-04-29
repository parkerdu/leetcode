package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	     Val int
	     Left *TreeNode
	    Right *TreeNode
 }



 // 解法一：递归方法
func isBalanced(root *TreeNode) bool {
	isBBT, _ := process(root)
	return isBBT
}


func process(node *TreeNode) (bool, int) {
	// step1: 判断node是否为空，空返回(true, 0)
	if node == nil {
		return true, 0
	}
	// step2: 获取左右子树 返回信息
	isLeft, lh := process(node.Left)
	isRight, rh := process(node.Right)
	// step3: 当前节点是否平衡 == 左右子树都是 && 高度差小于等于1
	if rh > lh {
		lh, rh = rh, lh
	}
	isBanlanced := (isLeft && isRight) && ((lh - rh) <= 1)
	// step4: 计算高度返回，较大者+1
	return isBanlanced, lh+1
}



// 解法二： 非递归，利用后序遍历的特性
/*
后序遍历：左右根，可以从底向上获取左右子树的高度
 */
func isBalanced1(root *TreeNode) bool {
	if root == nil { return  true}
	// step1: 定义1个stack和一个辅助栈helpStack,用于存放节点, 定义1个depth map用于存放每个节点的深度，并将root节点入stack栈
	stack, helpStack := make([]*TreeNode, 0, 48), make([]*TreeNode, 0, 48)
	depth := make(map[*TreeNode]int)
	stack = append(stack, root)
	// step2: while stack非空, 弹出一个元素到helpStack，左子树非空左子树进stack, 右子树非空，右子树进栈，左右都空depth记录深度
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		helpStack = append(helpStack, cur)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left == nil && cur.Right == nil {
			depth[cur] = 1
		}
	}
	// step3: while help非空，help出栈并开始计算左右子树高度差，并往depth中记录当前节点深度，出现高度差大于1，return false
	for len(helpStack) > 0 {
		cur := helpStack[len(helpStack) - 1]
		helpStack = helpStack[:len(helpStack) - 1]
		lh, rh := depth[cur.Left], depth[cur.Right]
		if rh >lh {
			lh, rh = rh, lh

		}
		if lh - rh > 1 {
			return false
		}
		depth[cur] = lh +1
	}
	return true
}