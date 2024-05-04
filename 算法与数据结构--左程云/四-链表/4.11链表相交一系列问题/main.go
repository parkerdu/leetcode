package main

import "fmt"

/*
问题：两个有环或者无环链表是否相交，相交返回第一个相交节点，否则返回nil
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、如何判断一个链表是否有环，有则返回第一个环节点，否则返回nil
func hasCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 走到这里说明两指针相交，一定有环存在了
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 2、无环单链表相交问题，
// 该函数可判断从[head1, tail1]和[head2, tail2]两个链表之间的交点
func interNodeWithoutCycle(head1, head2, tail1, tail2 *ListNode) *ListNode {
	// 两个无环单链表相交问题
	// step1:先让两个链表都走到最后一个节点，并记录长度
	cur1, cur2 := head1, head2
	var length int
	for cur1.Next != tail1 {
		cur1 = cur1.Next
		length ++
	}
	for cur2.Next != tail2 {
		cur2 = cur2.Next
		length --
	}
	// step2：若cur1==cur2说明相交，否则不相交
	if cur1 != cur2 {
		return nil
	}
	// step3: 若相交，则先让长链表走差值步
	// 调整一下，让长链表为cur1, 短链表为cur2
	if length < 0 {
		// 说明cur2是长链表
		cur1, cur2 = head2, head1
		length = -length
	} else {
		cur1, cur2 = head1, head2
	}

	for length != 0 {
		cur1 = cur1.Next
		length --
	}

	// step4: 各走一步往下对比
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}

// 我的写法，有bug，在对两个有环单链表的讨论有问题， 例如刚开始loop1==loop2, 此时可能为环外相交，而你直接返回loop2
func interNode(head1, head2 *ListNode) *ListNode {
	loop1, loop2 := hasCycle(head1), hasCycle(head2)
	if loop1 == nil && loop2 == nil {
		// 两个无环单链表相交问题
		// step1:先让两个链表都走到最后一个节点，并记录长度
		//cur1, cur2 := head1, head2
		//var length int
		//for cur1.Next != nil {
		//	cur1 = cur1.Next
		//	length ++
		//}
		//for cur2.Next != nil {
		//	cur2 = cur2.Next
		//	length --
		//}
		//// step2：若cur1==cur2说明相交，否则不相交
		//if cur1 != cur2 {
		//	return nil
		//}
		//// step3: 若相交，则先让长链表走差值步
		//// 调整一下，让长链表为cur1, 短链表为cur2
		//if length < 0 {
		//	// 说明cur2是长链表
		//	cur1, cur2 = head2, head1
		//	length = -length
		//} else {
		//	cur1, cur2 = head1, head2
		//}
		//
		//for length != 0 {
		//	cur1 = cur1.Next
		//  length --
		//}
		//
		//// step4: 各走一步往下对比
		//for cur1 != cur2 {
		//	cur1 = cur1.Next
		//	cur2 = cur2.Next
		//}
		//return cur1
		return interNodeWithoutCycle(head1, head2, nil,nil)
	} else if loop1 != nil && loop2 != nil {
		// 两个有环单链表相交问题
		// 1、环外相交，按无环单链表问题处理
		// 2、环内相交，返回loop1或者loop2皆可
		// step1: 判断在环内还是环外，让loop1一直往下走，若走一圈能遇到loop2说明环内，否则环外
		cur1 := loop1
		for {
			if cur1 == loop2 {
				return loop2 // bug处
			}
			cur1 = cur1.Next
			if cur1 == loop1 {
				break // 走完一圈也没有相遇，说明在环外相交
			}

		}
		// step2: 若是环内，返回loop1
		// step3: 若是环外，从head到loop节点为无环单链表相交问题
		return interNodeWithoutCycle(head1, head2, loop1, loop2)
	} else {
		// 一个链表有环，一个无环不可能相交
		return nil
	}
}

// 正确写法
func interNode1(head1, head2 *ListNode) *ListNode {
	loop1, loop2 := hasCycle(head1), hasCycle(head2)
	if loop1 == nil && loop2 == nil {
		return interNodeWithoutCycle(head1, head2, nil,nil)
	} else if loop1 != nil && loop2 != nil {
		// 两个有环单链表相交问题
		// 1、环外相交，按无环单链表问题处理
		// 2、环内相交，返回loop1或者loop2皆可
		// 3、不相交
		// step1: 判断是否为环外相交，if loop1 == loop2，可按环外处理
		if loop1 == loop2 {
			// 环外相交，从head到loop节点为无环单链表相交问题
			return interNodeWithoutCycle(head1, head2, loop1, loop2)
		} else {
			cur1 := loop1.Next
			for cur1 != loop1 {
				if cur1 == loop2 {
					return loop2
				}
				cur1 = cur1.Next
			}
			// 走到这说明不相交
			return nil
		}
		// step2: 若是环内，返回loop1
		// step3: 若是环外，从head到loop节点为无环单链表相交问题

	} else {
		// 一个链表有环，一个无环不可能相交
		return nil
	}
}

func main() {
	node5 := &ListNode{
		Val: 5,
		Next: nil,
	}
	node4 := &ListNode{
		Val: 4,
		Next: node5,
	}
	node3 := &ListNode{
		Val: 3,
		Next: node4,
	}
	node2 := &ListNode{
		Val: 2,
		Next: node3,
	}
	node5.Next = node3
	head1 := &ListNode{
		Val: 1,
		Next: node2,
	}

	node32 := &ListNode{
		Val:4,
		Next: node4,
	}
	node22 := &ListNode{
		Val: 3,
		Next: node32,
	}
	head2 := &ListNode{
		Val: 1,
		Next: node22,
	}

	node := interNode1(head1, head2)
	if node == nil {
		fmt.Println("不相交")
	} else {
		fmt.Println("相交，节点为：", node.Val)
	}
}