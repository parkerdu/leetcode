package main

import "fmt"

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

func (t *TreeNode) addLeft() {

}


const null = int(1) << 31

/*
给定一个层次遍历的数组生成二叉树
[1,
2,3,
4,nil, nil, 5,
nil, 6, nil, nil, nil, nil, 7, 8]
生成二叉树为
		          1
	      /   	         \
         2                  3
       /	\                \
	4          nil            5
   / \       /   \          /   \
nil   6    nil    nil      7	8
*/
func generateTree(queue []int) *TreeNode {
	cur := &TreeNode{}
	for val := range queue {
		if val != null {
			cur.Val = val

		}
	}
	return nil
}

func main() {
	fmt.Printf("%b", null)
}