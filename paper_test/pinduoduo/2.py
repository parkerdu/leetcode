import math
import heapq
# input = list(map(int,input().split(" ")))
# n,m,k = input[0],input[1],input[2]
n,m,k = 4,3,7

def pdd():
    data = []
    for i in range(1,m+1):
        for j in range(1,n+1):
            heapq.heappush(data, -i*j)
    for _ in range(k-1):
        heapq.heappop(data)
    print(-heapq.heappop(data))

def pdd4():
    temp = (8*k+1)/4
    t1 = -1/2+pow(temp,0.5)
    t = math.ceil(-1/2+pow(temp,0.5))
    if float(t) == t1:
        sub_num = t-1
        new_k = sub_num + 1
    else:
        # 不能被整处，往前走一个

        sub_num = t-1
        new_k = int(k - (t-1)*t/2)
    list = []
    index = []
    for i in range(sub_num+1):
        index.append([i,sub_num-i])
    for idx in index:
        list.append((n-idx[0])*(m-idx[1]))
    print(sorted(list,reverse=True)[new_k-1])
pdd()
pdd4()




