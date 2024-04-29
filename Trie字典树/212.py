# coding:utf-8
class Solution:
    """
    词搜索2
    法一：字典树加dfs(79题的升级）
    """
    def __init__(self):
        self.end_of_word = '#'
        self.root ={}

    def findWords(self, board, words):
        if not board or not words:
            return []
        # 应一个集合先存储结果再转成列表，这样做的目的就是去重
        # 因为同样一个单词可能在board中有两条以上的路可以走通，那就重复了
        results = set()
        for word in words:
            node = self.root
            for char in word:
                node = node.setdefault(char,{})
            node[self.end_of_word] = self.end_of_word
        print(self.root)

        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] in self.root:
                    result = self.dfs(board,i,j,self.root,'')
                    for r in result:
                        if r:
                            results.add(r)
        return list(results)

    def dfs(self, board, i, j, word, result):
        """我的问题代码"""
        # 这里我写的代码最后还剩一个问题就是，在字典树的同一个分支上有ab和abc都符合条件，但是由于aa满足self.end_of_word in word
        # 所以直接return，就把aac的情况丢掉了
        # 修改方法就是不让它return直接在函数外面定义一个变量，满足if条件就往里面加，不return就不会丢掉了
        if self.end_of_word in word:
            return [result]

        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] not in word:
            return []
        temp, board[i][j] = board[i][j], "@"
        # if board[i][j] in word:
        #     result += board[i][j]
        # 由于字典树的一条主分支中，可能存在多条子分支可以走通，所以最后的结果应该是所有结果加起来
        res = self.dfs(board, i-1, j, word[temp], result+temp) + self.dfs(board, i+1, j, word[temp], result+temp)\
               + self.dfs(board, i, j-1, word[temp], result+temp) + self.dfs(board, i, j+1, word[temp], result+temp)
        board[i][j] = temp
        return res



class Solution1:
    """解决我的问题代码"""
    def __init__(self):
        self.end_of_word = '#'
        self.root = {}
        self.results = set()

    def findWords(self, board, words):
        if not board or not words:
            return []
        # 应一个集合先存储结果再转成列表，这样做的目的就是去重
        # 因为同样一个单词可能在board中有两条以上的路可以走通，那就重复了
        for word in words:
            node = self.root
            for char in word:
                node = node.setdefault(char, {})
            node[self.end_of_word] = self.end_of_word
        print(self.root)

        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] in self.root:
                    self.dfs(board, i, j, self.root, '')
        return list(self.results)

    def dfs(self, board, i, j, word, result):
        """解决我的问题代码"""
        # 这里我写的代码最后还剩一个问题就是，在字典树的同一个分支上有ab和abc都符合条件，但是由于aa满足self.end_of_word in word
        # 所以直接return，就把aac的情况丢掉了
        # 修改方法就是不让它return直接在函数外面定义一个变量，满足if条件就往里面加，不return就不会丢掉了
        if self.end_of_word in word:
            self.results.add(result)

        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] not in word:
            return
        temp, board[i][j] = board[i][j], "@"
        # 由于字典树的一条主分支中，可能存在多条子分支可以走通，所以最后的结果应该是所有结果加起来
        self.dfs(board, i - 1, j, word[temp], result + temp)
        self.dfs(board, i + 1, j, word[temp],result + temp)
        self.dfs(board, i, j - 1, word[temp], result + temp)
        self.dfs(board, i, j + 1, word[temp],result + temp)
        board[i][j] = temp



class Solution2:
    """优化上面的四个重复代码，上下左右四连通图，只要定义一个for循环就可以了"""

    def __init__(self):
        self.end_of_word = '#'
        self.root = {}
        self.results = set()
        self.dx = [-1,1,0,0]
        self.dy = [0,0,-1,1]

    def findWords(self, board, words):
        if not board or not words:
            return []
        # 应一个集合先存储结果再转成列表，这样做的目的就是去重
        # 因为同样一个单词可能在board中有两条以上的路可以走通，那就重复了
        for word in words:
            node = self.root
            for char in word:
                node = node.setdefault(char, {})
            node[self.end_of_word] = self.end_of_word
        print(self.root)

        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] in self.root:
                    self.dfs(board, i, j, self.root, '')
        return list(self.results)


    def dfs(self, board, i, j, word, result):
        """优化上面的四个重复代码，上下左右四连通图，只要定义一个for循环就可以了"""
        if self.end_of_word in word:
            self.results.add(result)
        temp, board[i][j] = board[i][j], "@"
        for k in range(4):
            x = i + self.dx[k]
            y = j + self.dy[k]
            if 0 <= x < len(board) and 0 <= y < len(board[0]) and temp in word:
                self.dfs(board, x, y, word[temp], result + temp)
        board[i][j] = temp


if __name__ == "__main__":
    su = Solution2()
    board = [["a","b"],["c","d"]]
    # words =["aba","baa","bab","aaab","aaa","aaaa","aaba"]
    words = ["ab","cb","ac","bd","acd","abd"]
    # words = ["oath","pea","eat","rain"]
    print(su.findWords(board,words))
    print('a' or 'b')
