info = list(map(int, input().split(" ")))
n, m, k = info[0], info[1], info[2]
data = []
# n,m,k = 3,3,2
# data = [[2,3],[3,1]]
for i in range(k):
    abc = list(map(int,input().split(" ")))
    data.append(abc)

language = {}

for i in data:
    if i[1] in language:
        language[i[1]] += 1
    else:
        language[i[1]] = 1

print(n-max(language.values()))

