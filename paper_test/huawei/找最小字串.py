"""
input:a,b
out:a中的最小字串包含b
"""
def min_sub(s,t):
    """
    我的暴力解决方法  n = len(s) 时间O(N**2)
    :param s: 需要遍历的
    :param t: 目标字符串
    :return:
    """
    a, b = s, t
    length,l,r =  len(a),0,0,

    for i in range(len(a)):
        temp = list(b)
        if a[i] in temp:
            temp.remove(a[i])
            l = i
            for j in range(i+1,len(a)):
                if temp and a[j] in temp:
                    temp.remove(a[j])
                if not temp:
                    r = j
                    if r-l+1 < length:
                        length = r-l+1
                    break
    return length

class Solution:
    """
    leetcode 76 O(N)时间方法
    """
    def minWindow(self, s: str, t: str):
        if len(t) > len(s):
            return ''
        map_t = {}
        map_s = {}
        for i in t:
            map_t[i] = map_t.get(i,0) + 1
            map_s[i] = map_s.get(i,0) + 0
        l, r, count = 0,0,0
        res = ''
        while r < len(s):
            # 更新右指针
            if count < len(t):
                # 只有当s中的字符在t中时候，才会去更新map_s,count

                if s[r] in map_t:
                    map_s[s[r]] += 1
                    # 不仅r指针的字符在t中，还要满足个数小于等于t中的个数才会更新count
                    # 比如 s = bbb t = bb, s中的第三个b并不需要更新count
                    if map_s[s[r]] <= map_t[s[r]]:
                        count += 1

            # 匹配了就开始动左指针,当不匹配时候，退出while,继续走右指针
            while count == len(t):
                temp = s[l:r+1]
                if len(temp) < len(res) or not res:
                    res = temp
                if s[l] in map_t:
                    map_s[s[l]] -= 1
                    if map_s[s[l]] < map_t[s[l]]:
                        count -= 1
                l += 1
            r += 1
        return res



if __name__ == "__main__":
    a = 'abbdefg'
    b = 'bed'
    print(min_sub(a,b))
