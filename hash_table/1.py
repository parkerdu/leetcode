class Solution:
    def twoSum(self, nums, target):
        n = len(nums)
        dict_map = {}
        index = []
        for i in range(n):
            if nums[i] in dict_map:
                index.append(dict_map[nums[i]])
                index.append(i)
            else:
                dict_map[target-nums[i]] = i
            if len(index) == 2:
                return index

    def twoSum1(self, nums, target):
        dict_map = {}
        for i,x in enumerate(nums):
            if x in dict_map:
                return [i,dict_map[x]]
            else:
                dict_map[target-x] = i
