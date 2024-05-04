package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func commonPart(head1 *ListNode, head2 *ListNode) {
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			head1 = head1.Next
		} else if head1.Val > head2.Val {
			head2 = head2.Next
		} else {
			fmt.Println(head1.Val)
			head1, head2 = head1.Next, head2.Next
		}
	}
}




func commonPart2023(head1 *ListNode, head2 *ListNode) {
	for head1 != nil && head2 != nil {
		if head1.Val == head2.Val {
			fmt.Println(head1.Val)
			head1, head2 = head1.Next, head2.Next
		} else if head1.Val < head2.Val {
			head1 = head1.Next
		} else {
			head2 = head2.Next
		}
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
	head1 := &ListNode{
		Val: 1,
		Next: node2,
	}

	node32 := &ListNode{
		Val:4,
		Next: nil,
	}
	node22 := &ListNode{
		Val: 3,
		Next: node32,
	}
	head2 := &ListNode{
		Val: 1,
		Next: node22,
	}

	commonPart2023(head1, head2)
}
