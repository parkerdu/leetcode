class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return 0

        dp = [0 for _ in range(len(s)+1)]
        # dp[0] = 1 不是说一个空字符串有一个解，而是为了后面的递归的方便而设置的
        # 如字符串s='132'
        # i = 1时，dp[1] = 1
        # i =2 dp[2] = dp[1] + dp[0] = 1+1=2 这才是设置dp[0]=1的意义
        dp[0] = 1
        for i in range(1,len(s)+1):
            if i > 1:
                temp = s[i-2:i]
            # 字符串中第0个字符可以被单独解码
            if i == 1 and s[i-1] != '0':
                dp[i] = 1
            # 字符串中第i-1个字符可以被单独解码，单其和前面一个字符组成的两位数不能被单独解码，等于前一次的最优解
            # s[i - 1] != '0' 等价于 0<s[i-1]<10
            elif i > 1 and s[i-1] != '0' and (temp < '10' or temp > '26'):
                dp[i] = dp[i-1]
            # 字符串中第i-1个字符可以被单独解码，且其和前面一个字符组成的两位数也能被单独解码，等于前一次和上上一次的和
            elif i > 1 and s[i-1] != '0' and '09' < temp < '27':
                dp[i] = dp[i-1] + dp[i-2]
            # 字符串中第i-1个字符不可以被单独解码，但其和前面一个字符组成的两位数能被单独解码，于上上一次的最优解
            elif i > 1 and s[i-1] == '0' and '09' < temp < '27':
                dp[i] = dp[i-2]
        return dp[-1]


    def numDecodings1(self, s):
        """
        上面代码的简化版本
        因为上面的代码完全是按照动态规划的递推公式写出来的，观察其代码可以简化
        只要s[i-1]可以被单独解码就加上上一次的最优解
        只要s[i-2:]可以被单独解码就再加上上一次的最优解
        每次循环时候都检查这两个条件，发生了我就加上，不发生就不动
        :type s: str
        :rtype: int
        """
        if not s:
            return 0

        dp = [0 for _ in range(len(s)+1)]
        # dp[0] = 1 不是说一个空字符串有一个解，而是为了后面的递归的方便而设置的
        # 如字符串s='132'
        # i = 1时，dp[1] = 1
        # i =2 dp[2] = dp[1] + dp[0] = 1+1=2 这才是设置dp[0]=1的意义
        dp[0] = 1
        for i in range(1,len(s)+1):

            if i > 1:
                temp = s[i-2:i]
            if s[i-1] != '0':
                dp[i] += dp[i-1]
            if i > 1 and '09' < temp < '27':
                dp[i] += dp[i-2]
        return dp[-1]

if __name__ == "__main__":
    su = Solution()

    s = '101'
    print(su.numDecodings1(s))
