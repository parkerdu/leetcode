"""
题目描述：
https://blog.csdn.net/ten_sory/article/details/81135422
也就是求连通分支的个数
"""

def quick_find(edges):
    """
    :param edges: 二维数组 如[[0,1], [0, 4], [1, 2], [1, 3], [5, 6], [6, 7], [7, 5], [8, 9]]
    :return n: 连通分支的个数
    时间复杂度：给定m条边,n个顶点 时间O(MN)
    对每一条边，你都需要去循环遍历长度为n的根节点列表，总共m条边，M*N
    """
    # step1:统计总共有多少个顶点
    vertex = set()

    for edge in edges:
        if edge[0] not in vertex:
            vertex.add(edge[0])
        if edge[1] not in vertex:
            vertex.add(edge[1])
    # n为顶点数,相当于n个连通分支
    n = len(vertex)

    # step2:初始化根节点
    root = []
    for i in range(n):
        root.append(i)
    # 初始化的root = [0,1,2,3,...,n-1]

    # step3:每来一条边更新根节点
    for edge in edges:
        # edge = [p,q] 如果p,q的根节点不同的话就需要把,root[p] 更新为root[q]
        if root[edge[0]] != root[edge[1]]:
            # 先把要改变的根节点暂时存起来，要不然后面变了就不能找出其他相同的了
            temp = root[edge[0]]
            for i in range(len(root)):
                if root[i] == temp:
                    root[i] = root[edge[1]]

            # 这个if 条件成立说明p,q各自代表的分支是不连通的，执行完上面的for-loop相当于把两块不连通的区域添加一条边，变成连通的了
            # 所以连通分支个数减一
            n -= 1
    return n


def quick_union(edges):
    """

    :param edges:
    :return:
    """
    # step1:统计总共有多少个顶点
    vertex = set()
    for edge in edges:
        if edge[0] not in vertex:
            vertex.add(edge[0])
        if edge[1] not in vertex:
            vertex.add(edge[1])
    # n为顶点数,相当于n个连通分支
    n = len(vertex)

    # step2:初始化父亲节点 root中现在存的是父亲节点，要想知道根节点，要用find函数找
    root = []
    for i in range(n):
        root.append(i)
    # 初始化的root = [0,1,2,3,...,n-1]

    # step3:每来一条边更新根节点
    def find(p):
        """
        用来找到p的根节点，找到p=root[p]就是根节点
        :param p:
        :return:
        """
        while p != root[p]:
            p = root[p]
        return p
    for edge in edges:
        # edge[p,q]
        if find(edge[0]) != find(edge[1]):
            # p的根节点节点 挂 到 q的根节点上
            root[find(edge[0])] = find(edge[1])
            n -= 1
    return n








if __name__ == '__main__':

    # list = [
    #     [0,1],
    #     [0,4],
    #     [3,2],
    #     [2,4],
    #     [5,6],
    #
    # ]
    list = [
        (4, 3),
        (3, 8),
        (6, 5),
        (9, 4),
        (2, 1),
        (8, 9),
        (5, 0),
        (7, 2),
        (6, 1),
        (1, 0),
        (6, 7)
    ]

    print(quick_union(list))
