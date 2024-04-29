class Solution:
    """反转字符串"""
    def reverseString(self, s):
        """
        Do not return anything, modify s in-place instead.
        """

        left = 0
        right = len(s)-1
        while left <= right:
            s[left], s[right] = s[right], s[left]
            left += 1
            right -= 1

if __name__ == "__main__":
    a = ["h", "e", "l", "l", "o"]
    su = Solution()
    su.reverseString(a)
    print(a)