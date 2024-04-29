class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid):
        for i in range(len(obstacleGrid)):
            # row = 0
            # column = 0
            for j in range(len(obstacleGrid[0])):
                if obstacleGrid[i][j] == 1:
                    # if i == 0:
                    #     row += 1
                    # if j == 0:
                    #     column += 1
                    obstacleGrid[i][j] = 0
                elif i == 0 and j == 0:
                    obstacleGrid[i][j] = 1
                elif i == 0 and j >= 1 and obstacleGrid[i][j-1] != 0:
                    obstacleGrid[i][j] = 1
                elif j == 0 and i >= 1 and obstacleGrid[i-1][j] != 0:
                    obstacleGrid[i][j] = 1
                else:
                    if i != 0:
                        if j != 0:
                            obstacleGrid[i][j] = -1
        # print(obstacleGrid)
        for i in range(len(obstacleGrid)):
            for j in range(len(obstacleGrid[0])):
                if obstacleGrid[i][j] == -1:
                    obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]

        return obstacleGrid[-1][-1]
if __name__ == "__main__":
    su = Solution()
    a = [
  [0,1,0],
  [0,1,0],
  [0,0,0]
]
    print(su.uniquePathsWithObstacles(a))