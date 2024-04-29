class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def sortList(self, head):
        if not head or not head.next:
            return head
        mid = get_middle(head)
        left = head
        right = mid.next
        mid.next = None
        left_sorted = self.sortList(left)
        right_sorted = self.sortList(right)
        return merge(left_sorted, right_sorted)

def merge(left, right):
    dummy = ListNode(0)
    tail = dummy
    while left and right:
        if left.val < right.val:
            tail.next = left
            left = left.next
        else:
            tail.next = right
            right = right.next
        tail = tail.next
    tail.next = left if left else right
    return dummy.next

def get_middle(head):
    slow = fast = head
    while fast.next and fast.next.next:
        slow = slow.next
        fast = fast.next.next
    return slow

if __name__ == '__main__':
    node4 = ListNode(3)
    node3 = ListNode(2, node4)
    node2 = ListNode(1, node3)
    node1 = ListNode(4, node2)
    res = Solution().sortList(node1)
    while not res :
        print(res.val)
        res = res.next



