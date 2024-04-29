import math
# x = int(input())
x = 2
def find_zuhe(x):
    zuhe = []
    i = 1
    j = x-1
    while 1 <= i <= math.ceil(x/2) and j >= math.floor(x/2):
        zuhe.append([i,j])
        i += 1
        j -= 1
    return zuhe

def sougou1():
    res = {}
    for f3 in range(2,x+1):
        zuhes= find_zuhe(f3)
        for zuhe in zuhes:
            f1,f2 = zuhe[0],zuhe[1]
            idx = shulie(f1,f2,x)
            if idx >= 3:
                res[idx] = res.get(idx,0) + 1

    for key in sorted(res.keys()):
        print(key,end=' ')
        print(res[key])
        # print(key,'',res[key])



def shulie(f1,f2,x):
    f1, f2 = f1, f2
    count = 2
    while f2 <= x:
        f1,f2 = f2, f1+f2

        count += 1
        if f2 == x:
            return count
    return count
sougou1()

print('C'+' '+'8')
print('C'+' '+'3')
print('C'+' '+'3')
print('B'+' '+'0')
print('B'+' '+'0')
print('C'+' '+'0')
print('C'+' '+'0')

