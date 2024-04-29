def jingdong1(data):
    m, n = 5, 5
    for i in range(m):
        for j in range(n):
            # 计算连通数
            if data[i][j] != '#':
                num_of_grid,indexs = get_num(data,i,j)
                if num_of_grid >= 3:
                    # 消除连通
                    for idx in indexs:
                        data[idx[0]][idx[1]] = '#'
                    # 修复重力下落
                    rebuild_data = change_data(data)
                    data = rebuild_data
    res = 0
    for i in range(m):
        for j in range(n):
            if data[i][j] != '#':
                res += 1
    print(res)

def get_num(data,i,j,indexs):
    if [i,j] not in indexs:
        indexs.append([i,j])
    # 左边
    if i >=1 and data[i-1][j] == data[i][j]:
        get_num(data,i-1,j,indexs)
    if data[i][j+1] == data[i][j]:
        get_num(data,i,j+1,indexs)
    if data[i][j-1] == data[i][j]:
        get_num(data,i,j-1,indexs)
    if data[i+1][j] == data[i][j]:
        get_num(data,i+1,j,indexs)
    return indexs