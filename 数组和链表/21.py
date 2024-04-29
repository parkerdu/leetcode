# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):

    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        start = new = ListNode(1)
        p1 = l1
        # p1 = l1
        p2 = l2
        while p1 and p2:
            if p1.val <= p2.val:
                new.next = p1
                p1 = p1.next
            else:
                new.next = p2
                p2 = p2.next
            new = new.next
        if p1:
            new.next = p1
        if p2:
            new.next = p2
        return start.next


