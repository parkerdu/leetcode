class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        """所有是2的幂次方的数在二进制上都有一特点就是只有一个1
            如2**0
              2**1
              2**2"""
        if n != 0 and n & (n-1) == 0:
            return True
        else:
            return False

    def isPowerOfTwo(self, n: int) -> bool:
        """
        上面代码的简化
        :param n:
        :return:
        """
        return n > 0 and not n & (n-1)