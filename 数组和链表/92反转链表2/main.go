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

// 2024/04/30 22:58开始  23:28  用时30min oc, 严重说明熟练度不够
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var leftENd, rightHead, reverseHead, reverseTail *ListNode
	cur := head
	i := 0
	for cur != nil {
		i++
		if i < left {
			leftENd = cur
		} else if left <= i && i <= right {
			if reverseHead == nil {
				reverseTail = cur
			}
			next := cur.Next
			cur.Next = reverseHead
			reverseHead, cur = cur, next //这里逗号的左右会影响结果吗？
			continue
		} else {
			break
		}
		cur = cur.Next
	}
	// 重新链接
	if leftENd != nil {
		leftENd.Next = reverseHead
	}
	if reverseTail != nil {
		reverseTail.Next = rightHead
	}
	if leftENd == nil {
		return reverseHead
	}
	return head
}

// 优化减小1个变量
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var leftENd, reverseHead, reverseTail *ListNode
	cur := head
	i := 0
	for cur != nil {
		i++
		if i < left {
			leftENd = cur
		} else if left <= i && i <= right {
			if reverseHead == nil {
				reverseTail = cur
			}
			next := cur.Next
			cur.Next = reverseHead
			reverseHead, cur = cur, next //这里逗号的左右会影响结果吗？
			continue
		} else {
			break
		}
		cur = cur.Next
	}
	// 重新链接
	if reverseTail != nil {
		reverseTail.Next = cur
	}
	if leftENd != nil {
		leftENd.Next = reverseHead
	} else {
		return reverseHead
	}
	return head
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	// step1: 先遍历一遍链表，找到左边尾部，和右边头部，和中间链表头
	// step2: 反转中间链表
	// step3: 拼接
}
