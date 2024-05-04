//package main
//
//import (
//	"fmt"
//	"math"
//	"strconv"
//	"strings"
//)
//
///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//// 15:50 开始   18:00 首次做2h仍然没有ac
///*
//第一次做不出来原因分析，第一次针对树多的case会超时，
//因为想的是用满二叉树来做，这样对于有很多nil的节点，空间耗费会很大
//所以不能用满二叉树来做这道题，
//*/
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	if root == nil {
//		return ""
//	}
//	// step1: 求深度，depth
//	depth := deep(root)
//	// step2: 层次遍历，未超过depth的空置为#
//	res := ""
//	queue := make([]*TreeNode, 0)
//	// curDepth 当前深度，curLevelNum, nextLevelNum
//	var curDepth, curLevelNum, nextLevelNum int
//	queue = append(queue, root)
//	curDepth = 1
//	curLevelNum = 1
//	for len(queue) != 0 {
//		cur := queue[0]
//		queue = queue[1:]
//		// 我这种在出队列时候序列化
//		res += marshalInt(cur.Val) + "_"
//		if cur.Left == nil && curDepth+1 <= depth {
//			cur.Left = &TreeNode{Val: math.MaxInt}
//		}
//		if cur.Left != nil {
//			queue = append(queue, cur.Left)
//			nextLevelNum++
//		}
//
//		if cur.Right == nil && curDepth+1 <= depth {
//			cur.Right = &TreeNode{Val: math.MaxInt}
//		}
//		if cur.Right != nil {
//			queue = append(queue, cur.Right)
//			nextLevelNum++
//		}
//
//		//
//		curLevelNum--
//		if curLevelNum == 0 {
//			// 结算
//			curDepth++
//			curLevelNum, nextLevelNum = nextLevelNum, 0
//		}
//	}
//	return res
//}
//
//func marshalInt(v int) string {
//	if v == math.MaxInt {
//		return "#"
//	}
//	return strconv.FormatInt(int64(v), 10)
//
//}
//
//func unmarshal(s string) int {
//	if s == "#" {
//		return math.MaxInt
//	}
//	rootVal, _ := strconv.ParseInt(s, 10, 64)
//	return int(rootVal)
//}
//
//// 前序遍历求深度
//func deep(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	left := deep(root.Left)
//	right := deep(root.Right)
//	if left < right {
//		left = right
//	}
//	return 1 + left
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	if data == "" {
//		return nil
//	}
//	// step1: 分割字符串，得到节点数组
//	d := strings.Split(data, "_")
//	d = d[0 : len(d)-1]
//	//root := &TreeNode{Val: unmarshal(d[0])}
//	//// step2: 安装层次遍历，填充节点
//	//for i := range d {
//	//	if d[i] == "#" {
//	//		continue
//	//	}
//	//	if 2*i+1 < len(d) && d[2*i+1] != "#" {
//	//		queue[i].Left = queue[2*i+1]
//	//	}
//	//	if 2*i+2 < len(queue) && queue[2*i+2].Val != math.MaxInt {
//	//		queue[i].Right = queue[2*i+2]
//	//	}
//	//}
//	//return queue[0]
//	return nil
//}
//
///**
// * Your Codec object will be instantiated and called as such:
// * ser := Constructor();
// * deser := Constructor();
// * data := ser.serialize(root);
// * ans := deser.deserialize(data);
// */
//
//
//func main() {
//	// 构造，1,2,3，nil,nil,4,5
//	/*
//				 1
//			   2   3
//		          4  5
//	*/
//	six := &TreeNode{Val: 6}
//	five := &TreeNode{
//		Val: 5,
//	}
//	four := &TreeNode{Val: 4, Left: six}
//	three := &TreeNode{
//		Val:   3,
//		Left:  four,
//		Right: five,
//	}
//	two := &TreeNode{
//		Val: 2,
//	}
//	root := &TreeNode{Val: 1, Left: two, Right: three}
//	obj := Constructor()
//	s := obj.serialize(root)
//	fmt.Println(s)
//	five = obj.deserialize(s)
//}
