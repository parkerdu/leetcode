# coding:utf-8
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumNumbers(self, root):
        if not root:
            return 0
        self.result = 0
        self.help(root,root.val)
        return self.result

    def help(self,root,val):
        # 每到叶子节点，把值拿走
        # not(A OR B) = not A and not B
        if not (root.left or root.right):
            self.result += val
        if root.left:
            self.help(root.left, val*10+root.left.val)
        if root.right:
            self.help(root.right,val*10+root.right.val)

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
    tree.add(4)
    tree.add(9)
    tree.add(0)
    tree.add(5)



    su = Solution()
    print(su.sumNumbers(tree.root))