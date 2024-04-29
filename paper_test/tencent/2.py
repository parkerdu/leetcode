from math import factorial
from itertools import combinations
# list1 = list(map(int, input().split(" ")))
# t, k = list1[0],list1[1]
# data = []

# for i in range(t):
#     a = list(map(int, input().split(" ")))
#     data.append(a)
t, k = 4,89
data = [
    [90,10],
    [20,60],
    [400,800]
]

def cni3(n,i):
    result = 1
    for j in range(1, i+1):
        result = result * (n-i+j) // j
    return result


def tencent2():
    for idx in data:
        res = 0
        for l in range(idx[0],idx[1]+1):
            if l < k or k == 0:
                res += 1
            else:
                n = l//k

                for i in range(n+1):
                    b = l-k*i
                    res += cni3(i+b,i)
        print(res)
tencent2()


