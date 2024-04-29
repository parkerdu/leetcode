from copy import deepcopy
"""
题目描述：https://blog.csdn.net/u010507768/article/details/82590452 维克多博士核反应堆
"""

def dp(V, N, M, volumn, masses, energies):
    """
    :param V: 最大体积限制
    :param N: 液体个数
    :param M: 最大质量限制
    :param volumn: 体积列表
    :param masses: 质量列表
    :param energies: 能量列表
    :return: 最大能量值
    我的解法：创建一个全0的二维数组
    ·        该题是一个二位动态规划问题，画出来的更新图应该是一个立方体
                创建一个全0的(V+1)*(M+1)二维数组，让后更新N次这个数组就可以了
            1.定义子问题opt(i,v,m)为有{1,2,3,...i}个液体在体积限制为v,质量限制为m时候的最大能量值
            2.递归关系：if i = 0    opt(i,v,m) = 0
                        if vi>v or mi>m  肯定选不了第i瓶液体，只能为opt(i-1,v,m)
                        else:  既可以选第i瓶也可以不选第i个，这时候取两种情况的最大值 opt(i,v,m) = max(opt(i-1,v,m), ei+opt(i-1,v-vi,m-mi))
    """
    dp = []
    for j in range(V+1):
        temp1 = [0] * (M+1)
        dp.append(temp1)
    for i in range(N):
        # 这里重新开了一个新数组来暂时保存dp，如果不这样做的话，会导致错误，因为我当前层全部用的是上一层dp来更新的
        # 但是当v,m变大后再用的dP就很可能被你改变了
        # 例如更新dp[3][2]时候用到dp[1][2]了，这时候的dp[1][2]就不是上一层的了，；而是该层了
        # 更优化的办法就是让v,m从大到小变化，这样既不用开辟新空间，也不会影响结果，因为小的v,m更新不会用到大的v,m
        temp = deepcopy(dp)
        for v in range(1, V+1):
            for m in range(1,M+1):
                if volumn[i] > v or masses[i] > m:
                    pass
                else:
                    t1 = energies[i] + temp[v-volumn[i]][m-masses[i]]
                    t2 = temp[v][m]

                    dp[v][m] = max(energies[i] + temp[v-volumn[i]][m-masses[i]], temp[v][m])

    print(dp)
    return dp[-1][-1]

def dp1(V, N, M, volumn, masses, energies):
    """
    :param V: 最大体积限制
    :param N: 液体个数
    :param M: 最大质量限制
    :param volumn: 体积列表
    :param masses: 质量列表
    :param energies: 能量列表
    :return: 最大能量值
    优化版本
    """
    dp = []
    for j in range(V+1):
        temp1 = [0] * (M+1)
        dp.append(temp1)
    for i in range(N):
        for v in range(V, 0, -1):
            for m in range(M, 0, -1):
                if volumn[i] > v or masses[i] > m:
                    pass
                else:
                    dp[v][m] = max(energies[i] + dp[v-volumn[i]][m-masses[i]], dp[v][m])

    print(dp)
    return dp[-1][-1]
print(dp1(100,5,15,[50,40,30,20,10],[1,2,3,9,5],[300,480,270,200,180]))
print(dp1(5,3,3,[1,3,2],[1,2,3],[20,20,30]))