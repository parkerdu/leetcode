def shunfeng2(n, x, dp):
    l, r = 0, n-1
    while l <= r:
        mid = (l+r)//2
        if dp[mid-1] <= x < dp[mid]:
            return mid
        elif x < dp[mid-1]:
            r = mid - 1
        else:
            l = mid + 1
    return l

# n = [5]
# data = [1,2,1,3,4]
n = list(map(int, input().split(" ")))
data = list(map(lambda x: int(x), input().split()))
dp = [0]*(n[0]+1)
dp[0] = data[0]
length, j = 1, 0
for i in range(1, n[0]):
    if data[i] < dp[0]:
        j = 0
    elif data[i] >= dp[length-1]:
        j = length
        length += 1
    else:
        j = shunfeng2(length, data[i], dp)
    dp[j] = data[i]

print(length)