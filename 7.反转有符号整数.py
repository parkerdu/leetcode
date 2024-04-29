class Solution:
    def reverse(self, x: int) -> int:
        s_x = str(x)
        if s_x[0] == '-':
             res = -int(s_x[1:][::-1])

        else:
            res = int(s_x[::-1])
        if -2**31<=res<=2**31-1:
            return res
        return 0

    def reverse1(self,x):
        s = (x>0) - (x < 0)
        # 这里把 abs(x) 改成 s*x 时间更短
        r = int(str(abs(x))[::-1])
        return s*r*(r<2**31)

if __name__ == "__main__":
    su = Solution()
    print(su.reverse1(12300))



