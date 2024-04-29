# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        """我的解法，构造一个辅助函数，判断p,q是否在一个树里面，结果超时了"""
        while root:
            if self.helper(root.left,p,q) :
                root = root.left
            elif self.helper(root.right,p,q):
                root = root.right
            else:
                return root.val


    def helper(self,root,p,q):
        """辅助函数
            如果p,q都在以root为根节点的子树中，返回True
            否则，返回False"""
        if not root:
            return False
        stack = []
        is_p = False
        is_q = False
        cur = root
        while cur or stack:
            if cur is not None:
                stack.append(cur)
                cur = cur.left
            else:
                cur = stack.pop()
                # 因为p也是树的节点，所以要拿他的值
                if cur.val == p.val:
                    is_p = True
                if cur.val == q.val:
                    is_q = True
                cur = cur.right
        return is_p & is_q

    def lowestCommonAncestor1(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        # If looking for me, return myself
        if root == p or root == q:
            return root

        left = right = None
        # else look in left and right child
        if root.left:
            left = self.lowestCommonAncestor(root.left, p, q)
        if root.right:
            right = self.lowestCommonAncestor(root.right, p, q)

        # if both children returned a node, means
        # both p and q found so parent is LCA
        if left and right:
            return root
        else:
            # either one of the chidren returned a node, meaning either p or q found on left or right branch.
            # Example: assuming 'p' found in left child, right child returned 'None'. This means 'q' is
            # somewhere below node where 'p' was found we dont need to search all the way,
            # because in such scenarios, node where 'p' found is LCA
            # python的 or 返回第一个真值，都没有真值时返回最后一个表达式值
            # 所以在这里有三种情况：
                #1 left不是空 right is None-->返回left的值
                #2 left is None,right不是空-->返回right的值
                #3 两个都是None -->返回right的值,也即None
            return left or right

if __name__ == "__main__":
    class Node(object):
        def __init__(self, item):
            self.val = item
            self.left = None
            self.right = None


    class Tree(object):
        """二叉树"""

        def __init__(self):
            self.root = None

        def add(self, item):
            node = Node(item)
            if self.root is None:
                self.root = node
                return
            queue = [self.root]
            # if self.root=None  queue = [None] 此时列表不空，会进去wulie循环所以要在上面单独处理

            while queue:
                cur_node = queue.pop(0)
                # 为什么这里的cur_node 会有lchilld属性呢？
                # 因为你初始化的时候self.root = None 那么上面的if 条件会把node值给他，而node其实是一个类了
                if cur_node.left is None:
                    cur_node.left = node
                    return
                else:
                    queue.append(cur_node.left)
                if cur_node.right is None:
                    cur_node.right = node
                    return
                else:
                    queue.append(cur_node.right)

        def breadth_travel(self):
            """广度遍历"""
            if self.root is None:
                return
            queue = [self.root]
            while queue:
                cur_node = queue.pop(0)
                print(cur_node.val,end='')
                if cur_node.left is not None:
                    queue.append(cur_node.left)
                if cur_node.right is not None:
                    queue.append(cur_node.right)


    tree = Tree()
    tree.add(3)
    tree.add(5)
    tree.add(1)
    tree.add(6)
    tree.add(2)
    tree.add(0)
    tree.add(8)
    tree.add(9)
    tree.add(9)
    tree.add(7)
    tree.add(4)
    tree.breadth_travel()


    su = Solution()
    print('')
    print(su.lowestCommonAncestor1(tree.root,tree.root.left,tree.root.left.right.right))
    def su(a):
        return a or 4
    print(su(6))

