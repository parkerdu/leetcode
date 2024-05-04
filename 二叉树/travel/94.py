class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    """他给的列表是广度遍历的树
        实际上你可以推断他肯定是广度遍历，因为只有广度才可以给你就可以画出树，其他都需要两种遍历，还必须有中序
        [1,null,2,3] 第一行走完就1，第二行先是空再是2
        具体可以看visulize_tree代码运行结果"""

    def inorderTraversal(self, root):
        """非递归用栈实现中序遍历"""
        stack = []
        inorder = []
        cur = root
        # 直到cur是None,并且stack是空的时候结束循环，表示所有元素都遍历完了
        while cur or len(stack) > 0:
            # 先把左子树的所有左节点都压入栈,最后一个压入的就是最左边的叶子节点，待会他先出来
            if cur is not None:
                stack.append(cur)
                cur = cur.left
            else:
                cur = stack.pop()
                # 第一个出来的肯定是最左边的叶子节点，先把左节点打印，接下来是根，栈里面的下一个就是他的根
                inorder.append(cur.val)
                # 然后最左边叶子节点的右孩子肯定是None，所以再循环if不会执行，进入else弹出最左边叶子节点的根
                cur = cur.right
        return inorder

    def inorderTraversal1(self, root):
        """递归方法实现中序遍历"""
        if root is None:
            return []
        return self.inorderTraversal1(root.left) + [root.val] + self.inorderTraversal1(root.right)

