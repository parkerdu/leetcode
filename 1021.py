# coding:utf-8
import time
import datetime


class Solution:
    def removeOuterParentheses(self, S):
        # 我用了两个计数变量
#         left = 0
#         right = 0
#         s_remove = ''
#         j = 0
#         for i in range(len(S)):
#             if S[i] == '(':
#                 left += 1
#             if S[i] == ')':
#                 right += 1
#             if left == right:
#                 s1 = S[j:i+1]
#                 j = i+1
#                 s_remove += s1[1:len(s1)-1]

#         return s_remove
        # 别人只用一个count计数
        count = 0
        prev = None
        res = ''
        for i, c in enumerate(S):
            if count == 0:
                prev = i
            if c == '(':
                count += 1
            else:
                count -= 1
            if count == 0:
                res += S[prev + 1:i]
        return ''.join(res)

if __name__ == "__main__":
    t1 = datetime.datetime.now().microsecond
    # t3 = time.mktime(datetime.datetime.now().timetuple())
    S = "(()())(())(()(()))"
    su = Solution()
    print(su.removeOuterParentheses(S))
    t2 = datetime.datetime.now().microsecond
    # t4 = time.mktime(datetime.datetime.now().timetuple())
    strTime = 'funtion time use:%dms' % ((t2 - t1) * 1000 )
    print(strTime)
    # end = time.time()
    # print(str(end-start))
    # print('Running time: %s Seconds' % (end - start))
    print(bool(None))