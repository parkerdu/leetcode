package main

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

// 我的解法，时间O(n), 空间O(N), 借用栈来比较
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	stackA := make([]*ListNode, 0 ,16)
	stackB := make([]*ListNode, 0, 16)
	// 入栈
	for headA != nil {
		stackA = append(stackA, headA)
		headA = headA.Next
	}
	for headB != nil {
		stackB = append(stackB, headB)
		headB = headB.Next
	}

	// 出栈比较
	a, b := len(stackA)-1, len(stackB)-1
	var inter *ListNode
	for a >= 0 && b >= 0 {
		if stackA[a] == stackB[b] {
			inter = stackA[a]
		} else {
			return inter
		}
		a--
		b--
	}
	return inter
}

// 别人更好的方法：时间O(N), 空间O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var n1, n2 int
	cur := headA
	for cur != nil {
		n1++
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		n2 ++
		cur = cur.Next
	}

	// 将两个链表调整成相等长度
	for n1 != n2 {
		if n1 < n2 {
			// 则链表b长则，b往后走到
			headB = headB.Next
			n2--
		} else {
			headA = headA.Next
			n1--
		}
	}
	var inter *ListNode
	for headA != nil && headB != nil {
		if headA != headB {
			headA = headA.Next
			headB = headB.Next
		} else {
			inter = headA
			break
		}
	}
	return inter
}

// 不用计算长度，即可将长链表调整到和短链表相同
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b { // 当a和b一样长时候， 最后一起走到nil，循环结束
		if a == nil {
			a = headB // 如果a比b短的话，a先到nil, 此时让a等于更长的b,
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}


// 左程云老师代码
func getIntersectionNode(headA, headB *ListNode) *ListNode{
	cur1, cur2 := headA, headB
	var length int
	for cur1.Next != nil {
		cur1 = cur1.Next
		length ++
	}
	for cur2.Next != nil {
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
		cur1, cur2 = headB, headA
		length = -length
	} else {
		cur1, cur2 = headA, headB
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