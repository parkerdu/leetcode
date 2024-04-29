class Solution:
    def spiralOrder(self, matrix):
        """
        我的方法：用dfs去做的，写的特别麻烦，边界条件搞了半天
        :param matrix:
        :return:
        """
        if not matrix:
            return []
        m, n = len(matrix), len(matrix[0][:])
        i, j = 0, 0
        self.result = []
        while i < m and j < n:
            new_matrix = [row[j:n] for row in matrix[i:m]]
            x, y = m-i, n-j
            length = 0
            if x == 1:
                length = y
            elif y == 1:
                length = x
            else:
                length = 2*(x+y) - 4

            self.dfs(new_matrix,0,0,[new_matrix[0][0]],length)
            i += 1
            j += 1
            m -= 1
            n -= 1
        return self.result
    def dfs(self, board, i, j, temp,length):
        m, n = len(board), len(board[0][:])
        if len(temp) == length:
            self.result += temp

        if i == 0 and j < n-1:
            # 往右走
            temp += [board[i][j+1]]
            self.dfs(board, i, j+1, temp,length)
        if j == n-1 and i < m-1:
            # 往下走
            temp += [board[i+1][j]]
            self.dfs(board, i+1, j, temp,length)
        if i == m-1 and j > 0 and m != 1:
            # 往左走
            temp += [board[i][j-1]]
            self.dfs(board, i, j - 1, temp,length)
        if j == 0 and i > 1 and n != 1:
            # 往上走
            temp += [board[i-1][j]]
            self.dfs(board, i-1, j, temp,length)

    def spiralOrder1(self, matrix):
        """
        别人的方法，直接用循环的方法去做，跟我的想法是一模一样的，但是人家只用了12行代码就写出了
        注意：[[]] == true 因为 里面有一个空列表
        :param matrix:
        :return:
        """
        res = []
        while matrix:
            # 1.向右走
            # 顺时针走先把第0行按原来的顺序加进去
            res += matrix.pop(0)
            # 2.向下走
            # 接下来如果有第二行把第二行的最后一个元素加进去
            # 为啥这里需要matrix[0]条件因为[[]] == true 因为 里面有一个空列表
            if matrix and matrix[0]:
                for row in matrix:
                    res.append(row.pop())
            # 3.向左走
            # 接下来看有没有最后一行（如果输入只有一列，到这里就没有了，所以先判断一下）
            if matrix:
                # [::-1]为反转列表，把最后一行倒序添加
                res += matrix.pop()[::-1]
            # 4.向上走
            if matrix and matrix[0]:
                for row in matrix[::-1]:
                    res.append(row.pop(0))
        return res

if __name__ == "__main__":
    su = Solution()
    x = [
 [ 1 ],
 [ 4 ],
 [ 7 ]
]
    # x= [1,2,4]
    print(su.spiralOrder1(x))
    print(x[::-1])


