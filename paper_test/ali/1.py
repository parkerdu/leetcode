# list1 = list(input().split(','))
# list2 = list(input().split(','))
# x = int(input())


list1 = ['床前明月光','疑是地上霜','举头望明月','低头思故乡']
list2 = ['越光','床前地上霜','大地孤烟直','举头明月','低头是故乡']
x = 1
def Jaccrad(a, b):
    grams_reference = []
    grams_model = []
    for i in a:
        grams_reference.append(i)
    for i in b:
        grams_model.append(i)
    temp = 0
    for i in grams_reference:
        if i in grams_model:
            temp = temp + 1
    fenmu = len(grams_model) + len(grams_reference) - temp  # 并集
    jaccard_coefficient = float(temp / fenmu)  # 交集
    return jaccard_coefficient


def similarity():
    sims = []
    res = 0
    need_do = []
    for i in list1:
        temp = []
        for j in list2:
            sim = Jaccrad(i,j)
            temp.append(sim)
        max_sim = max(temp)
        idxs = []
        for idx,item in enumerate(temp):
            if item == max_sim:
                idxs.append(idx)
        idx = max(idxs)
        if idx in sims:
            repeat_idx = sims.index(idx)
            need_do.append(repeat_idx)
        res += max_sim
        sims.append(idx)

    for i in need_do:

    a = sims
    for i in range(len(a)):
        if i != len(a) - 1:
            print(a[i], end=',')
        else:
            print(a[i])



similarity()

