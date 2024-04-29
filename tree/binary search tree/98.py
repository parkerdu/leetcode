# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isValidBST(self, root) -> bool:
        """法一：中序遍历的二叉搜索树==一个升序的无重复元素的数组"""
        inorder = self.inorder(root)
        # 先弄成集合判重一下，因为如果有重复元素就不是BST,排序后再变成列表
        return inorder == list(sorted(set(inorder)))
    def inorder(self,root):
        if root is None:
            return []
        return self.inorder(root.left) + [root.val] + self.inorder(root.right)



    def isValidBST1(self, root) -> bool:
        """法一的优化,48ms击败99%
            优化1：用非递归方法实现中序遍历，比递归快
            优化2：在最后返回结果的时候，先用集合在排序O(nlogn)，再变集合，时间和空间消耗都是很大的
                    所以改成用O(N)时间来判断inorder是一个升序的无重复元素的数组"""
        if root is None:
            return True
        inorder = self.inorder1(root)
        for i in range(1,len(inorder)):
            if inorder[i] <= inorder[i-1]:
                return False
        return True

    def inorder1(self,root):
        stack = []
        res = []
        cur = root
        while cur or stack:
            if cur is not None:
                stack.append(cur.val)
                cur = cur.left
            else:
                cur = stack.pop()
                res.append(cur)
                cur = cur.right
        return res