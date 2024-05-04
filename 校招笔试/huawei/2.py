def is_prime(n):
    if n <= 1:
        return False
    i = 2
    while i < n:
        if n % i == 0:
            return False
        i += 1
    return True

def sum(a,b):
    prime = []
    for i in range(a,b):
        if is_prime(i):
            prime.append(i)
    if not prime:
        return 0
    else:
        decade = 0
        unit = 0
        for s in prime:
            s = str(s)
            if len(s) == 1:
                unit += int(s)
            else:
                d = int(s[-2])
                u = int(s[-1])
                decade += d
                unit += u
    print(min(decade,unit))

if __name__ == "__main__":
    nums = list(map(int,input().split(" ")))
    a, b = nums[0], nums[1]
    sum(a,b)