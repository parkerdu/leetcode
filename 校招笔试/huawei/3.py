start_name = input()
n = int(input())
all_name = []
name_set = set()

for i in range(n):
    line = input().split(',')
    all_name.append(set(line))

start_name = {start_name}


next_name = []
while True:

    flag = False
    while all_name:
        cur_name = all_name.pop()
        if start_name & cur_name:
            flag = True
            start_name = start_name | cur_name
        else:
            next_name.append(cur_name)
    if not flag:
        break
    all_name = next_name
print(len(start_name))