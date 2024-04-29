class Solution:
    def maxArea(self, height) -> int:
        left, right = 0, len(height)-1
        m = 0
        while left <= right:
            s = (right-left)*min(height[left],height[right])
            m = max(s,m)
            if height[left] <= height[right]:
                left += 1
            else:
                right -= 1
        return m

if __name__ == "__main__":
   a = [1, 8, 6, 2, 5, 4, 8, 3, 7]
   su = Solution()
   print(su.maxArea(a))