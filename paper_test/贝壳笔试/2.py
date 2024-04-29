import numpy as np
def missing_card(nums):
    """
    :param nums: 0-9出现的次数 len(nums)=10
    :return: n和x,n为总的卡片个数，x为缺失的卡片
    """
    n=1
    nums = np.array(nums)

    temp = np.array([0,1,0,0,0,0,0,0,0,0])
    while not (temp >= nums).all():
        n += 1
        for i in str(n):
            temp[int(i)] += 1
    comp = temp <= nums
    x_index = []
    for i,b in enumerate(comp):
        if not b:
            x_index.append(i)
    x_possilble = permuteUnique(x_index)
    x = []
    for row in x_possilble:
        a = [str(x) for x in row]
        temp = int("".join(a))
        if temp <= n:
            x.append(temp)
    for i in sorted(x):
        print('n', n)
        print('x', i)



def permuteUnique(nums):

    """
    46题的升级版，减枝就完事了
    对dfs中的for循环来说，遇到相同的数就不用再进递归了，出来的结果是一样的
    例如[1,1,2,3]
    第一个1进入下一层递归后，第二个1就没必要在做了，因为都是重复的
    :type nums: List[int]
    :rtype: List[List[int]]
    """

    result = dfs(nums, [], [])
    print(result)
    return result

def dfs(nums, path,res):
    if not nums:
        res.append(path)
    temp = set()
    for i in range(len(nums)):
        if nums[i] not in temp:
            dfs(nums[:i]+nums[i+1:],path+[nums[i]],res)
            temp.add(nums[i])
    return res

a = [2,12,9,3,3,3,3,2,2,2]
missing_card(a)

