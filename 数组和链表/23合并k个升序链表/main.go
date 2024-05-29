package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := (l + r) / 2
	left := merge(lists, l, mid)
	right := merge(lists, mid+1, r)

	// 合并left和right
	return mergeTwo(left, right)
}

func mergeTwo(left, right *ListNode) *ListNode {
	l, r := left, right
	root := &ListNode{}
	res := root
	for l != nil && r != nil {
		// step1: 左边小，
		if l.Val < r.Val {
			root.Next = l
			root = root.Next
			l = l.Next
		} else {
			root.Next = r
			root = root.Next
			r = r.Next
		}
	}
	if l != nil {
		root.Next = l
	}
	if r != nil {
		root.Next = r
	}
	return res.Next
}
