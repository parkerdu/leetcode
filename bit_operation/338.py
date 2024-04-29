class Solution:
    def countBits(self, num: int):
        """
        法一：用前面191题计算海明距离的方法来循环计算位数
        :param num:
        :return:
        """
        res = []
        for i in range(num+1):
            count = 0
            while i:
                i = i & (i-1)
                count += 1
            res.append(count)
        return res

    def countBits1(self, num: int):
        """
        法二：采用递推的方法
        先开一个数组来存储所有要计算位数的数
        如 res = [0,1,2,3]
        0的位数就是0，不用管
        1中1的个数=把1的最后一个1干掉之后1的个数 + 1
        2中1的个数=把2的最后一个1干掉之后1的个数 + 1
        3中1的个数=把3的最后一个1干掉之后1的个数 + 1
        因为一个数把1干掉之后肯定是减小了
        :param num:
        :return:
        """
        res = []
        for i in range(num+1):
            res.append(i)
        for i in range(1,num+1):
            # i & (i - 1) 为把i干掉一个1之后的十进制数
            # res[i&(i-1)] 为 这个十进制数中1的个数
            res[i] = res[i&(i-1)] + 1
        return res

    def countBits1(self, num: int):
        """
        法三：动态规划
        Index : 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15

        num :   0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4

        Do you find the pattern?

        Obviously, this is overlap sub problem, and we can come up the DP solution. For now, we need find the function to implement DP.

        dp[0] = 0;

        dp[1] = dp[0] + 1;

        dp[2] = dp[0] + 1;

        dp[3] = dp[1] +1;

        dp[4] = dp[0] + 1;

        dp[5] = dp[1] + 1;

        dp[6] = dp[2] + 1;

        dp[7] = dp[3] + 1;

        dp[8] = dp[0] + 1;
        ...

        This is the function we get, now we need find the other pattern for the function to get the general function. After we analyze the above function, we can get
        dp[0] = 0;

        dp[1] = dp[1-1] + 1;

        dp[2] = dp[2-2] + 1;

        dp[3] = dp[3-2] +1;

        dp[4] = dp[4-4] + 1;

        dp[5] = dp[5-4] + 1;

        dp[6] = dp[6-4] + 1;

        dp[7] = dp[7-4] + 1;

        dp[8] = dp[8-8] + 1;

        递推关系式
        dp[i] = dp[i-offset] +1
        0<i<2           offset = 1
        i没到2之前 2 =<i <4 offset = 2
        i没到2之前 4 =<i <8 offset = 4
        先开一个数组来存储所有要计算位数的数
        如 res = [0,1,2,3]
        """
