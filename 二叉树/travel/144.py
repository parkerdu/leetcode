# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:

    def preorderTraversal(self, root):
        """前序遍历递归解法"""
        if root is None:
            return []
        return [root.val]+self.preorderTraversal(root.left)+self.preorderTraversal(root.right)

    def preorderTraversal1(self, root):
        """前序遍历循环解法
            思想：用栈实现
            每出栈一个根节点，先入right,再入left
            则实现根>左>右
            用栈来控制循环，到了遍历的最后一个节点，就没有东西可以入栈了，这时候退出循环，也遍历完了所有元素"""
        if root is None:
            return []
        stack, res = [root], []
        while stack:
            cur = stack.pop()
            if cur:
                res.append(cur.val)
                stack.append(cur.right)
                stack.append(cur.left)
        return res





























