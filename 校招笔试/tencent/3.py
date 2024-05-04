input = list(map(int, input().split(" ")))
n,p,q = input[0],input[1],input[2]

# n,p,q = 2,1,0
def c(n,i):
    result = 1
    for j in range(1, i+1):
        result = result * (n-i+j) // j
    return result
num_p = []
cases = []
for m in range(p,n-q+1):
    num_p.append(m)
    cases.append(c(n,m))
total = sum(cases)

res = 0
for i in range(len(cases)):
    t = cases[i]/total * num_p[i]
    res += t
res = res % 1000000007
print(res)


