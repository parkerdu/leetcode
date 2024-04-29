class Solution(object):
    def permuteUnique(self, nums):

        """
        46题的升级版，减枝就完事了
        对dfs中的for循环来说，遇到相同的数就不用再进递归了，出来的结果是一样的
        例如[1,1,2,3]
        第一个1进入下一层递归后，第二个1就没必要在做了，因为都是重复的
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        self.result = []
        self.dfs(nums, [])
        return self.result

    def dfs(self, nums, path):
        if not nums:
            self.result.append(path)
        temp = set()
        for i in range(len(nums)):
            if nums[i] not in temp:
                self.dfs(nums[:i]+nums[i+1:],path+[nums[i]])
                temp.add(nums[i])


if __name__ == "__main__":
    a = [1,1,2]
    su = Solution()
    print(su.permuteUnique(a))