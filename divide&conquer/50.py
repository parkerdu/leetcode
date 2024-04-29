class Solution:
    def myPow(self, x: float, n: int) -> float:
        """递归解法
            思想：n奇数：x^n=x*x^((n-1)/2)*x^((n-1)/2)
                n偶数：x^n=x^(n/2)*x^(n/2)"""
        if 0 <= n < 1:
            return 1
        if n==-1:
            return 1/x
        # 如果n为偶数
        if not n%2:
            mid = self.myPow(x,n/2)
            return mid*mid
        # 当n为奇数
        if n%2:
            mid = self.myPow(x,(n-1)/2)
            return mid*mid*x

    def myPow1(self, x: float, n: int) -> float:
        """我的非递归解法(还有bug存在）
            思想：先定好y=x or 1/x
            然后在循环的时候，不用再判断奇偶，
            """
        if n > 0:
            y=x
        if n < 0:
            y=1/x
            n = -n
        if n == 0:
            return 1
        while n >= 2 or n <= -2:
            y = y*y
            # 如果n为偶数，如n=4,2,1当n=1时结束循环
            # 如果n为奇书，如n=5,2.5,1.25,当n=1.25时结束，后面判断一下如果while循环退出了n>1说明n是偶数
            n = n/2
        if n > 1:
            # 该条件满足表示n为正奇数再乘一个x
            y = y*x
        if n<-1:
            # 该条件满足表示n为负奇数再乘一个1/x
            y = y*(1/x)
        return y

    def myPow2(self, x: float, n: int) -> float:
        """老师的非递归解法
            思想：
            """
        if n < 0:
            x = 1/x
            n = -n
        pow = 1
        while n:
            # 这个if条件在n为奇数时成立，所以等价于：n%2 != 0
            if n & 1:
                pow *= x
            x *= x
            # 右移一位相当于int(n/2) int()取整:负数向上取整，正数向下取整
            n >>= 1
        return pow
if __name__ == "__main__":
    su = Solution()
    # print(su.myPow1(34.00515,-3))
    print(su.myPow2(2, 10))
    print(int(-1.5))
    print(pow(34.00515,-3))
    print(5&1)