class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) == 1:
            return s
        res = ''
        for i in range(len(s)):
            for j in range(len(s)-1,i,-1):
                if s[j] == s[i]:
                    if self.is_palindrome(s[i:j+1]):
                        if j-i+1 > len(res):
                            res = s[i:j+1]
                        break

        if not res and s:
            return s[0]
        return res

    def is_palindrome(self,s):
        i = 0
        j = len(s) - 1
        while i <= j:
            if s[i] == s[j]:
                i += 1
                j -= 1
                continue
            else:
                return False
        else:
            # 正常循环完成才会执行该步骤
            return True

    def longestPalindrome1(self, s: str) -> str:
        """
        法二：从回文串中间向两边扩散
        :param s:
        :return:
        """
        # 我的写法
        # res =''
        # for i in range(len(s)):
        #     odd = self.find_huiwen(s,i,i)
        #     if odd and len(odd) > len(res):
        #         res = odd
        #     even = self.find_huiwen(s,i,i+1)
        #     if even and len(even) > len(res):
        #         res = even
        # return res

        # 简化版本
        res = ''
        for i in range(len(s)):
            res = max(res,self.find_huiwen(s,i,i),self.find_huiwen(s,i,i+1),key=len)
        return res
    def find_huiwen(self,s,i,j):
        # 我的写法
        # l, r = i, i
        # while i >= 0 and j <= len(s) - 1:
        #     if s[i] == s[j]:
        #         l, r = i, j
        #         i -= 1
        #         j += 1
        #     else:
        #         break
        # return s[l:r+1]

        # 简化版本
        while i >= 0 and j <= len(s)-1 and s[i] == s[j]:
            i -= 1
            j += 1
        return s[i+1:j]


if __name__ =="__main__":
    su = Solution()
    a = ''
    print(su.longestPalindrome1(a))

