class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head):
        cur = head
        # 先找到最后一个节点o(n)时间
        while cur.next:
            new = cur.next
            new.next = cur
            cur = cur.next
        # 头节点指向最后一个节点
        head = cur
