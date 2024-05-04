package main

import "math/bits"

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

// 2024/04/30 23:57开始  00:21 oc
func rightSideView(root *TreeNode) []int {
	/*
		第一次的想法，有问题，左边可能还有元素，例如
		 1
		2  3
	  4
	*/
	if root == nil {
		return []int{}
	}
	res := make([]int, 0, 16)
	for root != nil {
		res = append(res, root.Val)
		if root.Right != nil {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return res
}

func rightSideView(root *TreeNode) []int {
	// 层次遍历，每次将该层最后一个元素存入数组
	// step1: 定义curLevelNum为当前层的节点数量，nextLevelNum为下一层节点数量，queue为队列
	if root == nil {
		return []int{}
	}
	curLevelNum, nextLevelNum := 1, 0
	queue := make([]*TreeNode, 0, 16)
	queue = append(queue, root)
	res := make([]int, 0)
	// step2: 队列非空，则出队一个元素，左，右入队同时计算nextlevelNum, 然后cur--
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextLevelNum++
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextLevelNum++
		}
		curLevelNum--
		if curLevelNum == 0 {
			// 结算当前层
			res = append(res, cur.Val)
			curLevelNum = nextLevelNum
			nextLevelNum = 0
		}
	}
	// step3: curlevelNum ==0 开始结算，重置cur,next 打印当前层最后元素
	return res
}

// 递归解法
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	view(root, 1, &res)
}

// 和前序遍历很像，但是有一点不同，这里是根右左遍历，先让右边界一直下去，这样先访问到的就是当前层的最后一个元素
func view(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}
	// 当前深度第一次访问
	if depth-1 == len(*res) {
		*res = append(*res, root.Val)
	}
	view(root.Right, depth+1, res)
	view(root.Left, depth+1, res)
}
