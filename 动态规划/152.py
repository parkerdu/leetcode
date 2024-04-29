class Solution(object):
    def maxProduct(self, nums):
        """
        法一：双指针 O（N**2)时间
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return
        res = float('-inf')
        for i in range(len(nums)):
            temp = 1
            for j in range(i,len(nums)):
                temp *= nums[j]
                if temp > res:
                    res = temp

        return res
if __name__ == "__main__":
    su = Solution()
    a = [-1]
    print(su.maxProduct(a))