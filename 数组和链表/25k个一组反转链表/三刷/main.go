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

// 14:46-15:09
func reverseKGroup(head *ListNode, k int) *ListNode {
	var start, end *ListNode
	for head != nil {
		rh, rt, next := reverseK(head, k)
		if end == nil {
			start, end = rh, rt
		} else {
			end.Next = rh
		}
		end = rt
		head = next
	}
	return start
}

// 从head开始翻转k个，不足则保持不动
// return: 翻转后的头，翻转后的尾巴，下一个头
func reverseK(head *ListNode, k int) (rh, rt, new *ListNode) {
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
	// 不足k个，翻转回来
	if k > 0 {
		var reverse *ListNode
		for rh != nil {
			next := rh.Next
			rh.Next = reverse
			rh, reverse = next, rh
		}
		return reverse, nil, nil
	}
	// 足够返回
	return rh, rt, head
}
