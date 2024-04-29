def get_data():
    data = []
    try:
        while True:
            n = list(map(int, input().split(" ")))
            for i in range(n[0]):
                abc = list(map(int,input().split(" ")))
                data.append(abc)
    except:
        pass
    return data

def yuanfudao2():
    data = get_data()
    for d in data:
        res = 0
        number, persons = d[0], d[1:]
        persons = sorted(persons)
        not_zeros = get_not_zero(persons)
        first_zero = persons.index(0)
        while not_zeros >= 3:
            # persons = persons[first_zero:first_zero+3]
            for i in persons[first_zero:first_zero+3]:
                i -= persons[not_zeros]
            res += persons[not_zeros]
        print(res)

def get_not_zero(persons):
    count = 0
    for i,num in enumerate(persons):
        if num != 0:
            count += 1

    return count

if __name__ == "__main__":
    yuanfudao2()