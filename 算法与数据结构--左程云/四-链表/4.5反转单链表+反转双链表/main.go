package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// 我的解法，空间复杂的O(N), 每次开辟了一个新的node
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	reverse := &ListNode{Val: head.Val}
	for head.Next != nil{
		node := &ListNode{
			Val: head.Next.Val,
			Next: reverse,
		}
		reverse = node
		head = head.Next
	}
	return reverse
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var reverse *ListNode
	for head != nil{
		temp := head.Next
		head.Next = reverse
		reverse = head
		head = temp
	}
	return reverse
}


// 空间复杂的O(1)解法， 整个过程只有指针在动，不涉及新内存开辟
func reverseList2(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil{
		cur := head
		head = head.Next
		cur.Next = reverse
		reverse = cur
	}
	return reverse
}


func test() {
	//node5 := &ListNode{
	//	Val: 5,
	//	Next: nil,
	//}
	//node4 := &ListNode{
	//	Val: 4,
	//	Next: node5,
	//}
	//node3 := &ListNode{
	//	Val: 3,
	//	Next: node4,
	//}
	//node2 := &ListNode{
	//	Val: 2,
	//	Next: node3,
	//}


	head := &ListNode{
		Val: 1,
		Next: nil,
	}
	res := reverseList2023(head)
	for res != nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}



/*
===================2023年10月31日写的单链表反转==============
很惊讶自己这次竟然一下子就写出来了
*/
func reverseList2023(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var reverse *ListNode
	for head != nil {
		tmp := head.Next  // step1: 临时保存head.next
		head.Next = reverse  // step2: head.next指向新链表
		reverse = head    // step3: 更新新链表
		head = tmp   // step4: 更新head,继续for-loop
	}
	return reverse
}


/*
===================双链表反转==============
*/
type DoubleListNode struct {
	Val int
	Next *DoubleListNode
	Prev *DoubleListNode
}

// 我的解法，空间复杂的O(N), 每次开辟了一个新的node
func reverseDoubleList(head *DoubleListNode) * DoubleListNode {
	var reverse *DoubleListNode
	for head != nil {
		temp := head.Next
		head.Next = reverse
		head.Prev = temp
		reverse = head
		head = temp
	}
	return reverse
}






func testReverseDoubleList() {
	node5 := &DoubleListNode{
		Val: 5,
		Next: nil,
	}
	node4 := &DoubleListNode{
		Val: 4,
		Next: node5,
	}
	node3 := &DoubleListNode{
		Val: 3,
		Next: node4,
	}
	node2 := &DoubleListNode{
		Val: 2,
		Next: node3,
	}
	head := &DoubleListNode{
		Val: 1,
		Next: node2,
	}

	head.Prev = nil
	node2.Prev = head
	node3.Prev = node2
	node4.Prev = node3
	node5.Prev = node4

	res := reverseDoubleList2023(head)
	for res != nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}




/*
===================2023年11月01日写的双链表反转==============
很惊讶自己这次竟然一下子就写出来了
*/
func reverseDoubleList2023(head *DoubleListNode) * DoubleListNode {
	var reverse *DoubleListNode
	for head != nil {
		tmp := head.Next
		head.Next = reverse
		head.Prev = tmp
		reverse = head
		head = tmp
	}
	return reverse
}

func main() {

	testReverseDoubleList() // 测试双链表

	test() // 测试单链表
}