class Solution:
    def searchInsert(self, nums, target: int) -> int:
        """我的解法"""
        if nums[-1] < target:
            return len(nums)
        for i, num in enumerate(nums):
            if num >= target:
                return i

    def searchInsert1(self, nums, target: int) -> int:
        """和上面结果一样，如果整个循环可以执行完成，就说明目标比最后一个数还大"""

        for i, num in enumerate(nums):
            if num >= target:
                return i
        return len(nums)

if __name__ == "__main__":
    a = [1,3,5,6]
    su = Solution()
    print(su.searchInsert(a,7))