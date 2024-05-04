package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	tmp := head
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			fast = fast.Next
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转slow指向的链表
	for slow != nil {
		slow, slow.Next, fast = slow.Next, fast, slow
	}

	tail := fast

	// 此时 fast = 1->2->3-nul
	for fast != nil {
		if head.Val != fast.Val {
			return false
		}
		fast = fast.Next
		head = head.Next
	}

	// 恢复反转的链表, 反转后 slow
	for tail != nil {
		tail, tail.Next, slow = tail.Next, slow, tail
	}

	head = tmp
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	return true
}




// 2023.11.01二刷
func isPalindrome2023(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	// step1: 快指针走完，偶数个fase走到nil停，奇数个fast走到最后一个元素
	fast, slow := head, head
	left := head

	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	// step2: 从slow位置开始做单链表反转, 并记录right位置
	var right *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = right
		right = slow
		slow = tmp
	}



	// step3: 从左右链表同时开始往下走，直到有一个为nil 或者 有元素不同停止
	last := right  // 先记录一下最后一个元素，后面还要恢复链表

	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left, right = left.Next, right.Next
	}



	// step4: 恢复链表，再记录一次right的位置，简单一个单链表反转即可
	var reverse *ListNode
	for last != nil {
		tmp := last.Next
		last.Next = reverse
		reverse = last
		last = tmp
	}

	// todo 2023年二刷时候在恢复链表时候没有做好，右边虽然恢复了，但是确实step5:挂到head上面去

	// step5: head 此时为 1->2->3-nil 要把他挂到reverse上
	for head.Next != nil {
		head = head.Next
	}
	head.Next = reverse

	return true

}




func main() {
	node5 := &ListNode{
		Val: 1,
		Next: nil,
	}
	node4 := &ListNode{
		Val: 2,
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
	head := &ListNode{
		Val: 1,
		Next: node2,
	}
	if isPalindrome2023(head) {
		fmt.Println("是回文链表")
	} else {
		fmt.Println("不是回文链表")
	}
}