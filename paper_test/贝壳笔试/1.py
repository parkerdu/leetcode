def GetResult(K):
    n = 1

    left = compute_left(n)
    right = K
    while left <= right:
        n += 1
        left = compute_left(n)
    return n


def compute_left(n):
    left = 0
    for i in range(1,n+1):
        left += 1/i

    return left


print(GetResult(2))