class Solution:
    """按奇偶排序"""
    def sortArrayByParityII(self, A):
        # n = len(A)
        # A = sorted(A,key=lambda x:x%2)
        # for i in range(int(n/2)):
        #     if i%2 != 0:
        #         A[i],A[n-1-i] = A[n-1-i],A[i]
        # return A

        # n = len(A)
        # for i in range(n):
        #     # 如果i+A[i]是奇数，且i是奇书，那么A[i]一定为偶数
        #     if (i+A[i])%2 != 0 and i%2 != 0:
        #         for j in range(i,n):
        #             if (j + A[j]) % 2 != 0 and j % 2 == 0:
        #                 A[i],A[j] = A[j],A[i]
        #                 break
        #     # 如果i+A[i]是奇数，且i是偶数，那么A[i]一定为奇数
        #     if (i+A[i])%2 != 0 and i%2 == 0:
        #         for j in range(i,n):
        #             if (j + A[j]) % 2 != 0 and j % 2 != 0:
        #                 A[i],A[j] = A[j],A[i]
        #                 break
        # return A
        """两个指针的方法"""
        n = len(A)
        i = 0
        j = 1
        while i<=n and j<=n:
            if A[i]%2 == 0:
                i += 2
            elif A[j]%2 != 0:
                j += 2
            else:
                A[i],A[j] = A[j],A[i]
                i += 2
                j += 2
        return A
if __name__ == "__main__":
    l = [1,2,3,5,3,6,4,6 ]
    s = Solution()
    sort = s.sortArrayByParityII(l)
    print(sort)
