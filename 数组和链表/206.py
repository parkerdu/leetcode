class Node:
    def __init__(self, x):
        self.elem = x
        self.next = None


class Solution:
    def reverseList1(self, head):
        """
        :param head: 传过来的head 就是链表中的第一个节点，head.item = 1, head.next = Node(2)
        :return: cur last跑到None，cur 刚好指向最后一个节点，所以返回它就行了
        """
        prev, cur, last = None, head, head.next
        while last:
            cur.next = prev
            prev, cur, last = cur, last, last.next
        cur.next = prev
        return cur

    def reverseList(self, head):
        cur = head
        prev = None
        while cur:
            #             这是之前我写的错误，
            #  prev,cur,cur.next = cur,cur.next,prev
            # 分析一下为什么错了
            # 第一步：先看等号右边，先准备好cur,cur.next,prev
            # 第二步：开始赋值，先把cur赋给prev,所以prev指向cur也即向后移动了一个节点
            #                  再把cur.next赋值给cur,所以cur也向后移动了一个节点
            #                  这时候你再把prev赋值给cur.next可是上面一步已经把cur往后移了一部，已经不是你想的把cur.next指向prev了
            # 所以要在cur移动之前把cur.next赋值

            # 等号左边cur.next在cur之前正确
            # cur.next = prev 完成链表中当前节点的next指向前一个节点
            # prev,cur = cur,cur.next 由于prev,cur都是我们定义的指针。她两不属于链表，只是起指示作用
            # 这一步就是把这两个指针同时向右移动一步，链表并没有任何改变
            cur.next,prev,cur = prev,cur,cur.next

            # 等号左边cur.next在cur之前正确
            # prev, cur.next, cur = cur, prev, cur.next
        return prev


if __name__ == "__main__":
    # 初始化一个 1->2->3->None 的列表
    # a 为头节点1，由a.elem = 1, a.next = Node(2) 所以传这个链表时候只要把a传过去就行了，a就是第一个节点
    a = Node(1)
    for i in range(2,4):
        # cur = a 表示cur也指向a这个链表了

        cur = a
        # 这时候cur指针往后面走，表示cur指向a链表的后面节点了
        while cur.next != None:
            cur = cur.next
        # cur.next 加了一个节点，就代表a链表的1加了一个节点
        cur.next = Node(i)

    # 打印上面链表中的值，所以遍历链表的程序也可以这样写
    cur = a
    while cur:
        print(cur.elem,end=' ')
        cur = cur.next

    su = Solution()
    # 这里只要把上面 链表的a拿过来就行了，a.elem = 1, a.next = Node(2)
    t = su.reverseList1(a)

    print('')
    while t:
        print(t.elem,end=' ')
        t = t.next
