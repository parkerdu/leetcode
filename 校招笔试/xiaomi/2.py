data = list(map(int,input().split()))
def find_max_sub(nums):
    temp, res = 0, 0
    i = 0
    for j in range(len(nums)):
        temp = temp + nums[j]
        if temp > res:
            # 只有temp > res这件事发生，才会去更新res
            res = temp

        else:
            # 如果 temp < res,这时候没必要更新l,r,但是一旦temp < 0
            # 代表当前[i,j]这个子数组的和小于0，那么这一段数组我直接不要了，更新i,temp
            if temp < 0:
                i = j + 1
                temp = 0
    print(res)
# a = [1,2,3]
find_max_sub(data)