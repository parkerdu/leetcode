class Solution:
    def minDepth(self, root) -> int:
        """思想：广度优先遍历-->把104题拿过来加一行代码就可以实现最小深度"""
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
                # 添加的地方
                # 如果node是叶子节点，直接返回深度，此时即为最小深度，最大深度需要遍历完成，最小不需要
                if not node.left and not node.right:
                    depth += 1
                    return depth
            depth += 1
        return depth

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
    print(su.minDepth(tree.root))