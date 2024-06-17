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

func reorderList(head *ListNode) {
	// step1: 快慢指针找到中点
	f, s := head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	// step2: 从中点开始翻转后面链表
	var r *ListNode
	for s != nil {
		next := s.Next
		s.Next, r = r, s
		s = next
	}
	l := head
	// step3: l向右，r向走走
	for l != r && l.Next != r {
		// step1: 记录front,behind, l和r同时走一步
		front, behind := l, r
		l, r = l.Next, r.Next
		// step2: 将behind插入到front和l之间
		front.Next = behind
		behind.Next = l
	}
}
