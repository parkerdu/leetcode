class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        self.result = []
        self.dfs(digits,len(digits)-1,1)
        return self.result

    def dfs(self, digits, index, plus):
        # 终止条件
        if index < 0:
            if plus == 1:
                self.result.insert(0, 1)
            return
        if digits[index] + plus == 10:
            self.result.insert(0,0)
            self.dfs(digits, index-1, 1)
        else:
            digits[index] += 1
            self.result = digits[:index+1] + self.result

    def plusOne1(self, digits):

        """
        非递归方法
        :type digits: List[int]
        :rtype: List[int]
        """
        n = len(digits)
        p = 1
        for i in range(n-1,-1,-1):
            sum = digits[i] + p
            p = sum // 10
            if p == 1:
                digits[i] = 0
            else:
                digits[i] += 1
                break
        if p == 1:
            digits = [1]+digits
        return digits



if __name__ == "__main__":
    a = [9, 9]
    su = Solution()
    print(su.plusOne1(a))