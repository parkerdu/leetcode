# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root):
        """我的解法：模仿广度优先遍历
            在维护的队列中，元素由原来的数变成一个列表，每一个列表存了一层的所有元素，用一个for循环遍历一层"""
        if not root:
            return []
        queue = [[root]]
        result = []
        while queue[0]:
            level = queue.pop(0)
            help_queue = []
            help_result = []
            # 完成一个for loop 表示一层元素遍历完毕，且下一层的所有儿子也都加入了help_queue中
            for i in level:
                # 存储当前层的元素
                help_result.append(i.val)
                # 两个if存储下一层的儿子
                if i.left:
                    help_queue.append(i.left)
                if i.right:
                    help_queue.append(i.right)

            queue.append(help_queue)
            result.append(help_result)
        return result

    def levelOrder1(self, root):
        """上一种解法的优化，队列中的元素依然和广度优先一样保持一样，只是用队列的长度来达到控制每一层都存入一个列表的目的"""
        if not root:
            return []
        queue = [root]
        result = []
        while queue:
            # level_size就表示当前层由几个元素，有几个元素就做几次for循环
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
        return result

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
    su =  Solution()
    print(su.levelOrder1(tree.root))