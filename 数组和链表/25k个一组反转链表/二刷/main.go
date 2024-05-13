package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 主函数里面，每次循环调动reversek, 然后拼接
	var rh, rt *ListNode
	for head != nil {
		start, stop, next := reverseK(head, k)
		head = next
		if rt == nil {
			rh, rt = start, stop
			continue
		}
		rt.Next = start
		rt = stop
	}
	return rh
}

// 从当前head开始，若有k个，返回反转后的头，尾巴，下一个
// 不足k个，返回当前的头，nil, 下一个，此时下一个肯定是nil会退出循环
func reverseK(head *ListNode, k int) (rh, rt, nextH *ListNode) {
	// 从head开始进程k个反转
	rt = head
	for head != nil {
		next := head.Next
		head.Next = rh
		rh, head = head, next
		k--
		if k == 0 {
			return rh, rt, head
		}
	}
	// 出来说明不足k个，head  == nil
	var reverse *ListNode
	for rh != nil {
		next := rh.Next
		rh.Next = reverse
		reverse, rh = rh, next
	}
	return reverse, nil, head
}
