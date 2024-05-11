package main

import "fmt"

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

// 下面是面试后，又整理了一波后的代码，还有2个问题存在，1个是主函数里面的rt没有更新，一个是head传递到函数里面重新赋值，外部是不改变的
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil {
//		return head
//	}
//	var rh, rt *ListNode
//	for head != nil {
//		// 这里有一个很基础但又很严重的问题，将head指针传递到函数里面去并且改变了，外面的head会不会变？
//		// 答案是看你操作什么：若你把head重写赋值，外面的head不会变化，但是若你操作head.val, head.next外面会跟着一起变化
//		// 解决方法：法一：将函数里面更改的head返回出来，法2把head指针的地址传递进去
//		left, right := reverseK(head, k)
//		if rt == nil {
//			rh, rt = left, right
//			continue
//		}
//		rt.Next = left
//	}
//	return rh
//}
//
//func reverseK(head *ListNode, k int) (rh, rt *ListNode) {
//	rt = head
//	for head != nil {
//		next := head.Next
//		head.Next = rh
//		rh, head = head, next
//		k--
//		if k == 0 {
//			return rh, rt
//		}
//	}
//	// 走到这里说明不足k个，再做一次反转，head==nil
//	var res *ListNode
//	for rh != nil {
//		next := rh.Next
//		rh.Next = res
//		res, rh = rh, next
//	}
//	return res, nil
//}

//// 按法1修改：将head作为返回值返回
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil {
//		return head
//	}
//	var rh, rt *ListNode
//	for head != nil {
//		// 这里有一个很基础但又很严重的问题，将head指针传递到函数里面去并且改变了，外面的head会不会变？
//		// 答案是看你操作什么：若你把head重写赋值，外面的head不会变化，但是若你操作head.val, head.next外面会跟着一起变化
//		// 解决方法：法一：将函数里面更改的head返回出来，法2把head指针的地址传递进去
//		next, left, right := reverseK(head, k)
//		head = next
//		if rt == nil {
//			rh, rt = left, right
//			continue
//		}
//		rt.Next, rt = left, right
//
//	}
//	return rh
//}
//
//func reverseK(head *ListNode, k int) (next, rh, rt *ListNode) {
//	rt = head
//	for head != nil {
//		next := head.Next
//		head.Next = rh
//		rh, head = head, next
//		k--
//		if k == 0 {
//			return head, rh, rt
//		}
//	}
//	// 走到这里说明不足k个，再做一次反转，head==nil
//	var res *ListNode
//	for rh != nil {
//		next := rh.Next
//		rh.Next = res
//		res, rh = rh, next
//	}
//	return nil, res, nil
//}

// 按法2修改：将head的地址传入
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var rh, rt *ListNode
	for head != nil {
		// 这里有一个很基础但又很严重的问题，将head指针传递到函数里面去并且改变了，外面的head会不会变？
		// 答案是看你操作什么：若你把head重写赋值，外面的head不会变化，但是若你操作head.val, head.next外面会跟着一起变化
		// 解决方法：法一：将函数里面更改的head返回出来，法2把head指针的地址传递进去
		left, right := reverseK(&head, k)
		if rt == nil {
			rh, rt = left, right
			continue
		}
		rt.Next, rt = left, right
	}
	return rh
}

func reverseK(head **ListNode, k int) (rh, rt *ListNode) {
	rt = *head
	for *head != nil {
		next := (*head).Next
		(*head).Next = rh
		rh, *head = *head, next
		k--
		if k == 0 {
			return rh, rt
		}
	}
	// 走到这里说明不足k个，再做一次反转，head==nil
	var res *ListNode
	for rh != nil {
		next := rh.Next
		rh.Next = res
		res, rh = rh, next
	}
	return res, nil
}
func main() {
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	l1 = reverseKGroup(l1, 2)
	fmt.Println(l1)
}
