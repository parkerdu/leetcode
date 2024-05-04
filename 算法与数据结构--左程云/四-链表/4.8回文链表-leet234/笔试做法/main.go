package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool{
	if head.Next == nil {
		return true
	}
	tmp := head
	// 准备一个栈
	stack := make([]int, 0, 16)
	for head != nil {
		// 所有元素出栈
		stack = append(stack, head.Val)
		head = head.Next
	}

	n := len(stack) -1
	for tmp != nil {
		if tmp.Val != stack[n] {
			return false
		}
		n--
		tmp = tmp.Next
	}
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
	if isPalindrome(head) {
		fmt.Println("是回文链表")
	} else {
		fmt.Println("不是回文链表")
	}
}
