package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 假设head= 1->2->3-nil


	old := head
	// 该for, 创建新建一个node，并将该node插入到原来节点之后
	// 1->new(1)->2->new(2)->3->new(3)-nil
	for old != nil {
		node := &Node{
			Val:  old.Val,
			Next: old.Next,
		}
		old.Next = node
		old = node.Next
	}

	// 为new节点的random指针赋值，newP.Random = old.Random.Next
	old = head
	for old != nil {
		newP := old.Next
		if old.Random != nil {
			newP.Random = old.Random.Next
		}
		// 这里假设原链表右n个元素，加上nil为n+1个，新建node后为2n+1个，肯定为奇数个，所以old.next.next不会报错
		old = old.Next.Next

	}

	// split, 将新节点分离出来
	old = head
	newNode := head.Next
	for old != nil {
		newP := old.Next
		old.Next = newP.Next
		if newP.Next != nil {
			newP.Next = newP.Next.Next
		}
		old = old.Next
	}
	return newNode
}
