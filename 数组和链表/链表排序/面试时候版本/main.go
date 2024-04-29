

package main

import (
"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}


/*
下面为面试时候写出的版本：
问题1：findMiddle函数在处理只有2个节点时候，始终返回nil，会出现找不到中点情况，一直死循环向下递归
解决方法：多一个last变量，记录slow的上一次指针位置，最后将last.next指向空。或者是在findmiddle时候，若fast.next.next为nil就停止，然后再向前一步
问题2： 58行代码是 与 不是 或
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// step1:找到中点，返回中点记为mid
	mid := findMiddle(head)
	// step2: 左边递归找 head开始 ， 右边 从mid开始
	left := sortList(head)
	right := sortList(mid)
	// step3： 合并2个有序链表
	return mergeList(left, right)
}


// 找到中点并返回，
func findMiddle(head *ListNode) *ListNode {
	fast, slow := head,head
	// 快指针2步，慢指针1步
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	res := slow.Next
	slow.Next = nil
	return res
}

func mergeList(left, right *ListNode) *ListNode {
	var root *ListNode
	for left != nil || right != nil {
		cur := &ListNode{}
		if left.Val < right.Val {
			cur.Val = left.Val
			left = left.Next
		} else {
			cur.Val = right.Val
			right = right.Next
		}
		root.Next = cur
	}
	if left != nil {
		root.Next = left
	}
	if right != nil {
		root.Next = right
	}
	return root.Next
}



func main() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d", a + b)
		}
	}
}