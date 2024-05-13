package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 16:27
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// step1: head走left-1步，找到left,
	var pre *ListNode
	reverse := head
	k := right - left + 1
	for left > 1 {
		pre = head
		head = head.Next
		left--
	}
	// step2: 开始反转r-l+1个，得到rh,rt
	var rh, rt *ListNode
	rt = head
	for head != nil {
		next := head.Next
		head.Next = rh
		rh, head = head, next
		k--
		if k == 0 {
			break
		}
	}
	// step3: 相连，left->rh, rt-head
	rt.Next = head
	if pre != nil {
		pre.Next = rh
		return reverse
	} else {
		return rh
	}
}

func main() {
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	l1 = reverseBetween(l1, 2, 4)
	fmt.Println(l1)
}
