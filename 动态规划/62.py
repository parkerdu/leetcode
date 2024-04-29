class Solution(object):
    def uniquePaths(self, m, n):
        """
        到终点的不同路径数
        :type m: int
        :type n: int
        :rtype: int
        """
        temp = []
        for _ in range(n):
            temp.append(1)
        for i in range(m - 1):
            for j in range(1,n):
                temp[j] = temp[j] + temp[j-1]
        return temp[n-1]
if __name__ == "__main__":
    su = Solution()
    print(su.uniquePaths(3,7))
