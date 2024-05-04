import re
def get_data():
    data = []
    n = 0
    try:
        while True:

            input_data = list(map(int, re.split(',|;', input())))
            data = input_data[:-1]
            n = input_data[-1]
    except:
        pass
    return data,n


def pdd1():
    data, n = get_data()
    # data = [555503,772867,756893,339138,399447,40662,859037,628085,855723,974471,599631,636153,581541,174761,948135,411485,554033,858627,402833,546467,574367,360461,566480,755523,222921,164287,420256,40043,977167,543295,944841,917125,331763,188173,353275,175757,950417,284578,617621,546561,913416,258741,260569,630583,252845]
    # n = 10
    even = []
    odd = []
    for i in data:
        # odd
        if i & 1 == 1:
            odd.append(i)
        else:
            even.append(i)
    res = sorted(even,reverse=True) + sorted(odd,reverse=True)
    for i in range(n):
        if i != n-1:
            print(res[i],end=',')
        else:
            print(res[i])

if __name__ == "__main__":
    pdd1()