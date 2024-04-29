"""
题目描述：https://www.cnblogs.com/liugl7/p/11391332.html
用并查集-quick_find 做就行了
把所有人的当作顶点，求出各个连通分支，然后找第一个发消息的人所在的分支中的顶点数，就是最终的答案
"""

def quick_find(first_name,edges):
    """

    :param first_name: 第一个发消息的人 Jack
    :param edges: 二维数组，表示所有群 [[Jack,Tom,Anny,Lucy],
                                        [Tom,Danny],
                                        [Jack,Lily]
                                        ]
    :return:
    """
    # 1.计算总共有多少人
    all_names = set()
    for edge in edges:
        for name in edge:
            if name not in all_names:
                all_names.add(name)

    n = len(all_names)
    # 2.初始化根节点
    root = {}
    for name in all_names:
        root[name] = name

    # 3.每来一条边，更新root,
    for edge in edges:
        first = edge[0]
        for name in edge[1:]:
            temp = root[first]
            if root[first] != root[name]:
                for key,value in root.items():
                    if root[key] == temp:
                        root[key] = root[name]
    first_name = root[first_name]
    res = 0
    for name in root.values():
        if name == first_name:
            res += 1
    return res

if __name__ == "__main__":
    first_name = 'Jack'
    edges = [['Jack','Tom','Anny','Lucy'],
            ['Tom','Danny'],
            ['Jack','Lily']
                    ]
    print(quick_find(first_name,edges))

