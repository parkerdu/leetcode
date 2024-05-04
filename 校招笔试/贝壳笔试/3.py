def min_time(data):
    n = data[0][0]
    time = data[1:]
    tool = 1
    res = 0
    next = 0
    for row in time:
        if next:
            if row[0] < row[1]:
                temp = row[0]
                tool = 0
            elif row[0] > row[1]:
                temp = row[1]
                tool = 1
            else:
                temp = row[0]
                next = 1
        else:
            if tool:
                if row[1] < row[0]+row[2]:
                    temp = row[1]
                elif row[1] > row[0]+row[2]:
                    temp = row[0]+row[2]
                    tool = 0
                else:
                    temp = row[1]
                    next = 1
            else:
                if row[0] < row[1] + row[2]:
                    temp = row[0]
                elif row[0] > row[1] + row[2]:
                    temp = row[1] + row[2]
                    tool = 1
                else:
                    temp = row[0]
                    next = 1
        res += temp
    return res


# def dfs(trees,tool,whole_trees,time):
#     if trees >= whole_trees:
a = [[3],
     [20,40,20],
     [10,4,25],
     [90,100,5]]
print(min_time(a))