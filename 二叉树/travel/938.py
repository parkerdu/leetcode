# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def rangeSumBST(self, root, L: int, R: int) -> int:
        """法一：先中序遍历完成再，找合适的数来求和
            时间：O(N)"""
        inorder_list = self.inorder(root)
        result = 0
        for i in inorder_list:
            if L <= i <= R:
                result += inorder_list[i]
        return result

    def inorder(self, root):
        if not root:
            return []
        return self.inorder(root.left)+[root.val]+self.inorder(root.right)


    def rangeSumBST1(self, root, L: int, R: int) -> int:
        """法二：遍历时候减枝，从而使整个树不用全部遍历
            时间：<=O(N)"""
        if not root:
            return 0
        if L <= root.val <= R:
            self.rangeSumBST1(root.left, L, R) + root.val + self.rangeSumBST1(root.right, L, R)
        elif root.val < L:
            self.rangeSumBST1(root.right, L, R)
        else:
            self.rangeSumBST1(root.left, L, R)

