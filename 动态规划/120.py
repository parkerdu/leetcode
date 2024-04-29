class Solution(object):
    def minimumTotal(self, triangle):
        """
        :type triangle: List[List[int]]
        :rtype: int
        """
        if len(triangle) == 1:
            return triangle[0][0]
        dp = triangle
        for i in range(len(triangle)-2,-1,-1):
            for j in range(len(dp[i])):
                dp[i][j] = dp[i][j] + min(dp[i+1][j],dp[i+1][j+1])
        return dp[0][0]

    def minimumTotal1(self, triangle):
        """
        :type triangle: List[List[int]]
        :rtype: int
        优化空间复杂度为O(N)
        """
        if len(triangle) == 1:
            return triangle[0][0]
        dp = triangle[-1]
        for i in range(len(triangle)-2,-1,-1):
            for j in range(len(triangle[i])):
                dp[j] = triangle[i][j] + min(dp[j],dp[j+1])
        return dp[0]
su = Solution()
a = [
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
print(su.minimumTotal1(a))
