package main

import (
	"fmt"
	"strconv"
	"strings"
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 层级遍历
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 10000)
	var l, r int
	queue[r] = root
	r++
	res := strconv.FormatInt(int64(root.Val), 10) + ","
	for l < r {
		cur := queue[l]
		l++
		if cur.Left != nil {
			// 在入队列时候做序列化
			res += strconv.FormatInt(int64(cur.Left.Val), 10) + ","
			queue[r] = cur.Left
			r++
		} else {
			res += "#,"
		}
		if cur.Right != nil {
			// 在入队列时候做序列化
			res += strconv.FormatInt(int64(cur.Right.Val), 10) + ","
			queue[r] = cur.Right
			r++
			//queue = append(queue, cur.Right)
		} else {
			res += "#,"
		}
	}
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	s := strings.Split(data, ",")
	queue := make([]*TreeNode, 0)
	rootVal, _ := strconv.ParseInt(s[0], 10, 64)
	root := &TreeNode{Val: int(rootVal)}
	queue = append(queue, root)
	i := 1
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		// 当前节点的左孩子是否存在，通过i找到左孩子的val
		cur.Left = unmarshal(s[i])
		i++
		cur.Right = unmarshal(s[i]) // 通过i+1找到右孩子
		i++
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return root
}

// s == # 返回 nil, 否则返回节点
func unmarshal(s string) *TreeNode {
	if s == "#" {
		return nil
	}
	rootVal, _ := strconv.ParseInt(s, 10, 64)
	return &TreeNode{Val: int(rootVal)}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
func main() {
	// 构造，1,2,3，nil,nil,4,5
	/*
				 1
			   2   3
		          4  5
	*/
	six := &TreeNode{Val: 6}
	five := &TreeNode{
		Val: 5,
	}
	four := &TreeNode{Val: 4, Left: six}
	three := &TreeNode{
		Val:   3,
		Left:  four,
		Right: five,
	}
	two := &TreeNode{
		Val: 2,
	}
	root := &TreeNode{Val: 1, Left: two, Right: three}
	obj := Constructor()
	s := obj.serialize(root)
	fmt.Println(s)
	five = obj.deserialize(s)
}
