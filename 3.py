class Solution:
    def lengthOfLongestSubstring(self, s: str):
        """
        我的解法 时间O(MN)
        N = len(s)
        M = max(len(temp))
        :param s:
        :return:
        """
        temp = []
        length = 0
        for i in s:
            if i not in temp:
                temp.append(i)
            else:
                index = temp.index(i)
                temp = temp[index+1:]
                temp.append(i)
            length = max(length, len(temp))
        return length

    def lengthOfLongestSubstring1(self, s: str):
        """
        优化时间复杂度
        时间O(N)
        把列表换成字典，查询时间O(1)
        :param s:
        :return:
        """
        dict, length, start = {}, 0, 0
        for i, str in enumerate(s):
            if str in dict:
                index = dict[str]
                length = max(length, i-start)
                start = max(index + 1, start)
            dict[str] = i
        # 最后返回的时候要注意，如abcaefg,你前面程序跑完只可以得到length = 3,
        # 因为每从出现重复的才更新，而后面的bcaefg到最后都没重复，所以要再次更新一下length
        return max(length,len(s)-start)

if __name__ == "__main__":
    s = "abcadcbb"
    su = Solution()
    print(su.lengthOfLongestSubstring1(s))

