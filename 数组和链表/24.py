class Solution:

    def __init__(self, x):
            self.val = x
            self.next = None

    def swapPairs(self, head):
        """我的迭代解法，用了两个指针，情况讨论特别麻烦"""
        # 特殊1空链表
        if head is None:
            return head
        # 特殊2只有一个节点
        if head.next is None:
            return head
        cur = head
        head = cur.next
        cur.next = head.next
        head.next = cur
        prev, cur = cur, cur.next
        # 特殊2只有两个节点
        if cur is None:
            return head
        while cur.next:
            prev.next = cur.next
            cur.next = prev.next.next
            prev.next.next = cur
            prev, cur = cur, cur.next
            if cur is None:
                break
        return head

    def swapPairs1(self, head):
        """递归解法，用了三个指针，思路很清晰，也不需要做什么讨论情况
            每次递归完成两个交换，直到所有交换完成或者最后还剩一个节点不需要交换了，返回head"""
        # 当head为空或head只有一个节点时候不需要交换返回head
        if not head or not head.next:
            return head

        first, second = head, head.next
        third = second.next
        head = second
        second.next = first
        first.next = self.swapPairs(third)

        return head

    def swapPairs2(self, head):
        pre, pre.next = self, head
        while pre.next and pre.next.next:
            a = pre.next
            b = a.next
            pre.next, b.next, a.next = b, a, b.next
            pre = a
        return self.next


if __name__ == "__main__":
    pass


