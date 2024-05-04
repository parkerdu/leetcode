"""
甲刚开始有x的钱。每过一天，甲的钱会增加kx + 1这么多。问需要多少天，甲的钱会超过
10000000000000000。(0 <= x <= 1000000000000000000) (0 <= k <= 100)
输入描述
输入只有一行，包括两个整数x, k。
输出描述
输出只有一行，包括一个整数，表示所需天数。
输入样例一
0 1
输出样例一
54
输入样例二
9999999999990000 0
输出样例二
10000
"""
def how_many_days(x,k,limit):
    """

    :param x: 刚开始钱数目
    :param k: 参数
    :param limit: 10000000000000000
    :return:
    """
    money = x
    day = 0
    while money <= limit:
        money = (k+1)*money + 1
        day += 1
    return day

limit = 10000000000000000
print(how_many_days(9999999999990000,0,limit))
