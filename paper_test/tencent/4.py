# coding: utf-8
s_len = input()
s = input()
m = input()
count = 0
for i in range(int(m)):
    s1 = input()
    s1_len = len(s1)
    start_s, start_s1 = 0, 0
    while start_s < s_len:
        if s[start_s] != s1[start_s1]:
            break
        start_s, start_s1 = start_s+1, (start_s1+1) % s1_len
    else:
        count += 1
print(count)