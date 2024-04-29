class Solution(object):
    def get_max_min(self,a):
        if len(a) == 1:
            return a[0]
        n = len(a)
        mid_index = n // 2
        # n 奇数
        if n & 1:
            mi = min(self.get_max_min(a[:mid_index]),self.get_max_min(a[mid_index+1:]),a[mid_index])
            ma = max(self.get_max_min(a[:mid_index]),self.get_max_min(a[mid_index+1:]),a[mid_index])
        else:
            mi = min(self.get_max_min(a[:mid_index]),self.get_max_min(a[mid_index+1:]))
            ma = max(self.get_max_min(a[:mid_index]), self.get_max_min(a[mid_index + 1:]))
        return mi

if __name__ == "__main__":
    a = [1,3,6,8]
    su = Solution()
    print(su.get_max_min(a))