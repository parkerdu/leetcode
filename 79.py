class Solution:
    def exist(self, board, word: str) -> bool:
        for i in range(len(board)):
            for j in range(len(board[0])):
                # 如果是找到起点开始进入dfs
                # if board[i][j] == word[0]:
                #     if self.dfs(board, i, j, word):
                #         return True
                # 上面两个if可以合并
                if board[i][j] == word[0] and self.dfs(board, i, j, word):
                    return True
        return False

    def dfs(self, board, i, j, word):
        # 终止条件，如果可以把word遍历完成，说明可以走通，所有最多递归len(word)
        if len(word) == 0:
            return True
        # 两个限制，超过边界和起点不匹配
        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] != word[0]:
            return False

        # 把当前遍历的字母换成#就可以避免走重复的，因为如果它走到这个#上面会由于起点不匹配return False
        temp = board[i][j]
        board[i][j] = '#'
        res = self.dfs(board, i - 1, j, word[1:]) or self.dfs(board, i + 1, j, word[1:]) \
              or self.dfs(board, i, j - 1,word[1:]) or self.dfs(board, i, j + 1, word[1:])
        # 回退时候再把原来的board改回去
        board[i][j] = temp
        return res