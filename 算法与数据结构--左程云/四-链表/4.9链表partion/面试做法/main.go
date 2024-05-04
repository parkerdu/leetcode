package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func partionList(head *ListNode, pivot int) *ListNode {
	var sh, st, eh, et, bh, bt *ListNode
	for head != nil {
		if head.Val < pivot {
			if st == nil {
				sh, st = head, head
			} else {
				st.Next = head
				st = st.Next
			}
		} else if head.Val == pivot {
			if et == nil {
				eh, et = head, head
			} else {
				et.Next = head
				et = et.Next
			}
		} else {
			if bt == nil {
				bh, bt = head, head
			} else {
				bt.Next = head
				bt = bt.Next
			}
		}
		head = head.Next
	}

	if st != nil {
		st.Next = nil
	}
	if et != nil {
		et.Next = nil
	}
	if bt != nil {
		bt.Next = nil
	}

	if sh == nil && st == nil {
		if eh == nil && et == nil {
			if bh != nil && bt != nil {
				return bh
			} else {
				return nil
			}
		} else {
			if bh != nil && bt != nil {
				et.Next = bh
				return eh
			} else {
				return eh
			}
		}
	} else {
		if eh == nil && et == nil {
			if bh != nil && bt != nil {
				st.Next = bh
				return sh
			} else {
				return sh
			}

		} else {
			st.Next = eh
			if bh == nil && bt == nil {
				return sh
			} else {
				et.Next = bh
				return sh
			}
		}
	}
}

func partionList1(head *ListNode, pivot int) *ListNode {
	var sh, st, eh, et, bh, bt *ListNode
	for head != nil {
		if head.Val < pivot {
			if st == nil {
				sh, st = head, head
			} else {
				st.Next = head
				st = st.Next
			}
		} else if head.Val == pivot {
			if et == nil {
				eh, et = head, head
			} else {
				et.Next = head
				et = et.Next
			}
		} else {
			if bt == nil {
				bh, bt = head, head
			} else {
				bt.Next = head
				bt = bt.Next
			}
		}
		head = head.Next
	}

	if st != nil {
		st.Next = nil
	}
	if et != nil {
		et.Next = nil
	}
	if bt != nil {
		bt.Next = nil
	}

	// todo 找出和老师写的一样简单的讨论
	if sh == nil && st == nil {
		if eh == nil && et == nil {
			if bh != nil && bt != nil {
				return bh
			} else {
				return nil
			}
		} else {
			if bh != nil && bt != nil {
				et.Next = bh
				return eh
			} else {
				return eh
			}
		}
	} else {
		if eh == nil && et == nil {
			if bh != nil && bt != nil {
				st.Next = bh
				return sh
			} else {
				return sh
			}

		} else {
			st.Next = eh
			if bh == nil && bt == nil {
				return sh
			} else {
				et.Next = bh
				return sh
			}
		}
	}
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
	part := partionList(head, 2)
	for part != nil {
		fmt.Println(part.Val)
		part = part.Next
	}
}