# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        """思想：对于一个BST来说，p,q有三种情况：
            情况1：都在左子树中，那就再去左子树中找
            情况2：都在右子树中，去右子树中找
            情况3：除了1和2，the others we called them the others
            """
        while root:
            # 情况1：p,q都在左子树中
            if p.val < root.val and q.val < root.val:
                root = root.left
            # 情况2：p,q都在右子树中
            elif p.val > root.val and q.val > root.val:
                root = root.right
            # 情况3：上面两种都不满足时，要么p，q分别再左右子树中，要么p或者q就是root,所以剩下的情况的最近公共祖先就是root
            else:
                return root

    def lowestCommonAncestor1(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        """递归方法实现"""
        if p.val < root.val and q.val < root.val:
            self.lowestCommonAncestor1(root.left)
        if p.val > root.val and q.val > root.val:
            self.lowestCommonAncestor1(root.right)
        else:
            return root
