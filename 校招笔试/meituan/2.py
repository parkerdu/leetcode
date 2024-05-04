import sys
input("input:")
def get_data():
    data = []
    out = []
    try:
        while True:
            n = list(map(int, input().split(" ")))

            for i in range(n[0]):
                abc = input().split(" ")
                data.append(abc)
            for line in sys.stdin:
                a = list(map(int, line().split(" ")))
                out.append(a)
    except:
        pass
    return data,out

def max_length():
    data, out = get_data()
    for index in out:
        try:
            a = data[index[0]-1]
            b = data[index[1]-1]
            res = compute(a,b)

        except:
            pass
        else:
            print(res)

def compute(a,b):
    if len(a) > len(b):
        return compute(b,a)
    count = 0
    for i,s in enumerate(a):
        if a[i] == b[i]:
            count += 1
        else:
            break
    return count


print(max_length())