
from collections import defaultdict
n, m, k = list(map(lambda x: int(x), input().split()))

edges = []

tmp_dict = defaultdict(list)

person = [-1]+[0]*n
language = [-1]+[0]*m

for i in range(k):
    u, v = list(map(lambda x: int(x),input().split()))
    edges.append([u, v])

    tmp_dict[v].append(u)
    person[u] = 1

all_sets = []
for p in tmp_dict.values():
    all_sets.append(set(p))

cur_sets = []
for s in all_sets:
    new_cur_set = []
    s_c = [s]
    for c in cur_sets:
        if len(s & c):
            s_c.append(c)
        else:
            new_cur_set.append(c)

    s_c_n = set()
    if len(s_c) != 0:
        for ne in s_c:
            s_c_n |= ne
        new_cur_set.append(s_c_n)
    cur_sets = []
    for new_cur in new_cur_set:
        cur_sets.append(new_cur)

no_language = 0
for i in range(1, n+1):
    if person[i] == 0:
        no_language += 1

print(len(cur_sets)-1 + no_language)