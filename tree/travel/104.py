# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def maxDepth1(self, root) -> int:
        """思想：BFS(广度优先）完全照搬102题的层次遍历方法"""
        if not root:
            return 0
        queue = [root]
        result = []
        while queue:
            # level_size就表示当前层由几个元素
            level_size = len(queue)
            help_result = []
            # 不再需要辅助队列存储下一层的所有儿子了，现在所有儿子还是进入原先的队列，只不过我用一层的长度来控制输出
            for i in range(level_size):
                node = queue.pop(0)
                help_result.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            result.append(help_result)
        return len(result)

    def maxDepth(self, root) -> int:
        """BFS(广度优先)稍微修改上面的层次遍历"""
        if not root:
            return 0
        queue = [root]
        depth = 0
        while queue:
            # level_size就表示当前层由几个元素
            level_size = len(queue)
            # 不再需要辅助队列存储下一层的所有儿子了，现在所有儿子还是进入原先的队列，只不过我用一层的长度来控制输出
            for i in range(level_size):
                node = queue.pop(0)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            depth += 1
        return depth

    def maxDepth3(self, root):
        """DFS(深度优先)，分治"""
        if not root:
            return 0
        return 1+max(self.maxDepth3(root.left),self.maxDepth3(root.right))




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
                print(cur_node.val, end='')
                if cur_node.left is not None:
                    queue.append(cur_node.left)
                if cur_node.right is not None:
                    queue.append(cur_node.right)


    tree = Tree()
    tree.add(3)
    tree.add(9)
    tree.add(20)
    tree.add(6)
    tree.add(5)
    tree.add(15)
    tree.add(7)
    tree.add(4)
    su =  Solution()
    print(su.maxDepth3(tree.root))