class Solution:
    """生成所有有效的括号"""
    def generateParenthesis(self, n: int):
        self.list = []
        self._gen(0, 0, n, "")
        return self.list

    def _gen(self, left, right, n, result):
        if left == n and right == n:
            self.list.append(result)
            return
        if left < n:
            self._gen(left+1, right, n, result + "(")
        if left > right and right < n:
            self._gen(left, right+1, n, result + ")")









    def generateParenthesis1(self, n):
        def generate(p, left, right, parens=[]):
            if left:         generate(p + '(', left - 1, right)
            if right > left: generate(p + ')', left, right - 1)
            if not right:    parens += p,
            return parens

        return generate('', n, n)
if __name__ == "__main__":
    su = Solution()
    print(su.generateParenthesis(3))

