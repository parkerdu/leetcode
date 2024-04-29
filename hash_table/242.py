
class Solution:

    def isAnagram(self, s: str, t: str) -> bool:
        """法一：来一个删除一个
                """
        dict_1 = []
        dict_2 = []
        for i in s:
            dict_1.append(i)
        for i in t:
            dict_2.append(i)
        for i in range(len(dict_1)):
            if dict_1[i] in dict_2:
                dict_2.remove(dict_1[i])
            else:
                return False
        return not dict_2

    def isAnagram2(self, s: str, t: str) -> bool:
        """法二：hashmap方法
            把所有字母和个数存入字典，如果两个字典相同，则是异位词"""
        dict_1 = {}
        dict_2 = {}
        for i in s:
            # 第一次遇见字典get的方法，避免写if，dict_1.get(i,0)
            dict_1[i] = dict_1.get(i,0)+1
            # if i in dict_1:
            #     dict_1[i] += 1
            # else:
            #     dict_1[i] = 1
        for i in t:
            if i in dict_2:
                dict_2[i] += 1
            else:
                dict_2[i] = 1
        return dict_1 == dict_2

    def isAnagram3(self, s: str, t: str) -> bool:
        """法二同一种思想，老师的代码还是比我简洁"""
        dict_1, dict_2 = {}, {}
        for i in s:
            # dict.get(key,default) 该函数返回指定键的值，如果没有该键则返回默认值，这里默认值
            # 设置为0，没有就返回0
            dict_1[i] = dict_1.get(i,0)+1
        for i in t:
            dict_2[i] = dict_2.get(i,0)+1
        return dict_1 == dict_2

if __name__ == "__main__":
    s = "anagram"

    t = "nagaram"
    su = Solution()
    print(su.isAnagram3(s,t))