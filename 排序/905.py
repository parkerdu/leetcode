class Solution:
    """按奇偶排序"""
    def sortArrayByParity(self, A):
        n = len(A)
        first = 0
        last = n-1
        while first <= last:
            # first是偶数向右移动
            if not A[first] & 1:
                first += 1
            # last是奇数向左移动
            elif A[last] & 1:
                last -= 1
            else:
                A[first], A[last] = A[last], A[first]
                first += 1
                last -= 1
        return A


if __name__ == "__main__":
    a = [5]
    su = Solution()
    print(su.sortArrayByParity(a))
