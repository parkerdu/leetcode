

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// step1:找到中点，返回中点记为mid
	mid := findMiddle(head)
	// step2: 左边递归找 head开始 ， 右边 从mid开始
	left := sortList(head)
	right := sortList(mid)
	// step3： 合并2个有序链表
	return mergeList(left, right)
}




// 找到中点并返回，
func findMiddle(head *ListNode) *ListNode {
	fast, slow := head,head
	// 快指针2步，慢指针1步
	var last *ListNode
	for fast != nil && fast.Next != nil {
		last = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if last != nil {
		last.Next = nil
	}
	return slow
}

func mergeList(left, right *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return root.Next
}



func main() {
	l4 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  2,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  1,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  4,
		Next: l2,
	}
	res := sortList(l1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}