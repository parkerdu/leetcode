
board = []
for i in range(5):
    line = list(input().split())
    board.append(line)


def dfs(board, pre, cur_i, cur_j, cur_count):
    for di in range(-1, 2):
        for dj in range(-1, 2):
            if di != 0 or dj != 0:
                if di == 0 or dj == 0:
                    x = cur_i + di
                    y = cur_j + dj
                    if x >= 0 and x < 5 and y >= 0 and y < 5 and board[x][y] == pre:
                        tmp = board[x][y]
                        board[x][y] = 'x'
                        cur_count[0] += 1
                        dfs(board, pre, x, y, cur_count)
                        if cur_count[0] < 3:
                            board[x][y] = tmp


def down(board):
    for i in range(3, -1, -1):
        for j in range(0, 5):
            if board[i][j] != 'x':
                tmp_i = i
                while tmp_i+1 < 5 and board[tmp_i+1][j] == 'x':
                    board[tmp_i+1][j], board[tmp_i][j] = board[tmp_i][j], board[tmp_i+1][j]
                    tmp_i += 1

for i in range(5):
    for j in range(5):
        if board[i][j] != 'x':
            cur_count = [0]
            tmp = board[i][j]
            board[i][j] = 'x'
            dfs(board, tmp, i, j, cur_count)
            if cur_count[0] >= 3:
                down(board)
            else:
                board[i][j] = tmp

res = 0
for i in range(5):
    for j in range(5):
        if board[i][j] != 'x':
            res += 1

print(res)

