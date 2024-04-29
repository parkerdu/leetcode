# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None
class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        思想：判重，每过一个节点把他的地址存起来，因为节点值可能相同，地址是不同的
        列表插入元素是O(n)时间，因为后面元素的下标都要变化
        当然append是O(1)时间
        判断元素是否在列表中需要O(n)时间
        所以整个时间复杂度就是O(n**2)
        """
        cur = head
        l = []
        while cur:
            l.append(cur)
            cur = cur.next
            if cur.val in l:
                return True
        return False

    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        优化一下上面那种方法
        集合添加一个元素时间复杂度O(1)
        判断元素是否在集合中时间复杂度O(1)
        所以整体时间复杂度为O(n)
        """
        cur = head
        l = set()
        while cur:
            l.add(cur)
            cur = cur.next
            if cur and cur in l:
                return True
            if cur is None:
                return False
        return False

    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        第三种方法：运动会长跑
        在学校环形跑道上参加长跑，跑的快的人和跑得慢的人，会出现相遇
        那如果有环，快指针和慢指针也会相遇
        """
        fast = slow = head
        while slow:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True
            if slow is None or fast is None:
                return False
