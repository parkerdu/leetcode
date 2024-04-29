from functools import reduce
from operator import xor
class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return 2 * sum(set(nums)) - sum(nums)

    def singleNumber1(self, nums):
        """
        法二：按位异或，一个数，和同一个数按位异或两次不改变该数，因为相同的数异或的结果是0 ，所以拿0和所有数异或后，相当于和单独的一个数异或
        :type nums: List[int]
        :rtype: int
        """
        for i in range(1, len(nums)):
            nums[0] ^= nums[i]
        return nums[0]

    def singleNumber2(self, nums):
        """
        思想和法二是一样的，只是精简代码量而已
        :type nums: List[int]
        :rtype: int
        """
        return reduce(xor, nums)
if __name__ =="__main__":
    a = [4, 1, 2, 1, 2]
    su = Solution()
    print(su.singleNumber1(a))