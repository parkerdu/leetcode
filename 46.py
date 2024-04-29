# coding:utf-8
class Solution(object):

    def permute(self, nums):
        """
        法一：dfs方法
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        self.res = []
        self.dfs(nums, [])
        return self.res

    def dfs(self, nums, path):
        if not nums:
            self.res.append(path)
        for i in range(len(nums)):
            self.dfs(nums[:i] + nums[i+1:], path + [nums[i]])

    def permute1(self, nums):

        """
        法二：非递归方法
        :type nums: List[int]
        :rtype: List[List[int]]
        """

        def swap(i, j, nums):
            new_nums = list(nums)
            new_nums[i], new_nums[j] = new_nums[j], new_nums[i]
            return new_nums

        result = [nums, ]
        for i in range(len(nums) - 1):
            for one in result[:]:
                for j in range(i + 1, len(nums)):
                    result.append(swap(i, j, one))

        return result

if __name__ == "__main__":
    # a = [1,2,3]
    a = []
    for i in range(1,999):
        a.append(i)
    su = Solution()
    print(su.permute1(a))