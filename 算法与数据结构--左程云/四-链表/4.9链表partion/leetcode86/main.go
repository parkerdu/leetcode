package main

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

func partition(head *ListNode, x int) *ListNode {
	cur := head
	var sh, st, eh, et *ListNode
	// 分离为两个链表
	for cur != nil {
		if cur.Val < x {
			if st == nil {
				sh, st = cur, cur
			} else {
				st.Next = cur
				st = st.Next
			}
		} else {
			if et == nil {
				eh, et = cur, cur
			} else {
				et.Next = cur
				et = et.Next
			}
		}
		cur = cur.Next
	}

	// 合并成新链表
	if st != nil {
		st.Next = eh
		if et == nil {
			et = st
		}
	} else {
		sh = eh
	}
	if et != nil {
		et.Next = nil
	}
	return sh
}
