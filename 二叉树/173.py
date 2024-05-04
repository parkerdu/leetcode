# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Tree(object):
    """二叉树"""
    def __init__(self):
        self.root = None
    def add(self, item):

        node = TreeNode(item)
        if self.root is None:
            self.root = node
            return
        queue = [self.root]
        while queue:
            cur_node = queue.pop(0)
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


class BSTIterator(object):

    def __init__(self, root):
        """
        :type root: TreeNode
        """
        self.nodeStack = []
        self.leftBranchIntoStack(root)

    def leftBranchIntoStack(self, node):

        while (node):
            self.nodeStack.append(node)
            node = node.left

    def next(self):
        """
        @return the next smallest number
        :rtype: int
        """
        node = self.nodeStack.pop()
        self.leftBranchIntoStack(node.right)
        return node.val

    def hasNext(self):
        """
        @return whether we have a next smallest number
        :rtype: bool
        """
        if self.nodeStack:
            return True
        else:
            return False
# class BSTIterator(object):
#
#     def __init__(self, root):
#         self.stack = []
#         while root:
#             self.stack.append(root)
#             root = root.left
#
#     def next(self):
#         node = self.stack.pop()
#         r = node.right
#         while r:
#             self.stack.append(r)
#             r = r.left
#         return node.val
#
#     def hasNext(self):
#         return len(self.stack) != 0
class BSTIterator:

    def __init__(self, root):
        self.stack = []
        self.inorder(root)


    def inorder(self,node):
        if node is None:
            return

        self.inorder(node.left)
        self.stack.append(node.val)
        self.inorder(node.right)

    def next(self):
        """
        @return the next smallest number
        """
        next = self.stack.pop(0)
        return next

    def hasNext(self):
        """
        @return whether we have a next smallest number
        """
        if self.stack == []:
            return False
        else:
            return True


if __name__ == "__main__":
    root = Tree()
    root.add(15)
    root.add(2)
    root.add(16)
    # root.add(4)
    # root.add(5)
    print(root.root)
    obj = BSTIterator(root.root)
    param_1 = obj.next()
    print(param_1)
    print(obj.next())
    param_2 = obj.hasNext()
    print(param_2)