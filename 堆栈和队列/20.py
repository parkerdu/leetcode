class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        我自己开始没想出来怎么做，下面是老师的算法
        用栈取解决问题，首先栈的特性是先进后出，那么根据这种特性你想用它
        step1:找出问题中是否存在两个相反顺序的东西，本题中如果是一个合法的字符串，左括号和右括号刚好是两个相反的顺序
            例如'([{}]) 左括号是123，右括号是321
        step2:找到这两种相反的关系后，你就可以把一半压入栈，另一种相反顺序用来出栈
        """
        d = {'(': ')', '[': ']', '{': '}'}
        left = set(d.keys())
        right = set(d.values())
        s = s.lower()
        stack = []
        for i in s:
            if i in left:
                stack.append(i)
            elif i in right and len(stack) == 0:
                return False
            elif i in right and d[stack[-1]] is i:
                stack.pop(-1)
            # 如果i在右边但是不匹配前面的括号，那么一定不合法了，比如((]最后一个中括号就是这种情况
            else:
                return False
        if len(stack) == 0:
            return True
        else:
            return False

    def isValid1(self, s):
        """
        :type s: str
        :rtype: bool
        同样一种思想我自己写出来的代码用了20行，还很复杂的讨论了几种情况
        老师只用了9行代码，还不用讨论很多的复杂情况

        时间复杂度：O(n)
        遍历s需要O(n)
        查找是否在字典中需要O（1）
        列表append 需要O（1）
        所以总时间 O(n)

        空间复杂度：O（n)
        最坏的时候都是左括号，全都入栈了"""
        stack = []
        d = {')':'(',']':'[','}':'{'}
        for c in s:
            # 如果c not in d表示c 不在d.keys中，也就是c是左括号，压入栈
            if c not in d:
                stack.append(c)
            # 不满足上面条件，才会执行这个，那c一定是右括号，右括号的时候如果stack是空的就一定不是有效的，如()]
            # 或者是右括号但是不匹配，那也是错的例如((]
            elif not stack or d[c] != stack.pop():  # pop默认剪切最后一个元素
                return False
        # stack 空列表时候布尔值为False,not 取反就是true
        # 所以对not stack  空列表：True
        #                  非空：False
        return not stack


if __name__ == "__main__":
    s = ""
    su = Solution()
    print(su.isValid(s))

    # 测试判断字符是否在字典中，是对键还是值
    d = {'key':'value'}
    if 'key' in d:
        print('for key')
    if 'value' in d:
        print('for value')
    # 结果是对键操作

    # 排序也是对键
    d = {3:'a',1:'b',2:'c'}

    print(sorted(d))

    stack = []
    print(not stack)

    stack = [1,3,5]
    print(stack.pop())