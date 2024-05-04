def get_data():
    data = []
    max_time = 0
    try:
        while True:
            n = list(map(int, input().split(" ")))
            max_time = n[1]
            data = list(map(int,input().split(" ")))

    except:
        pass
    return max_time, data

def yuanfudao1():
    max_time, data = get_data()
    # max_time, data = 2,[4,3,3,3,1,5,5]
    dict_data = {}
    for i in data:
        if i in dict_data:
            dict_data[i] += 1
        else:
            dict_data[i] = 1
    for key, v in dict_data.items():
        if v > max_time:
            dict_data[key] = '#'
    for i in data:
        if dict_data.get(i) != '#':
            print(i, end=' ')


if __name__ == "__main__":
    yuanfudao1()