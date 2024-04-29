class Solution:
    def majorityElement(self, nums) -> int:
        dict_fre = {}
        for i in nums:
            if i in dict_fre:
                dict_fre[i] += 1
            else:
                dict_fre[i] = 1
        for i in dict_fre.keys():
            if dict_fre[i] > int(len(nums)/2):
                return i

    def majorityElement1(self, nums) -> int:
        """优化版本，这是我第二次遇到这种写法了，第一次教了你，你在这里并没有学会用"""
        dict_fre = {}
        for i in nums:
            # dict_fre.get(i,0)的意思时如果字典中有i就把i的value拿出来，没有默认为0
            dict_fre[i] = dict_fre.get(i,0) + 1
        for i in dict_fre.keys():
            if dict_fre[i] > int(len(nums)/2):
                return i

if __name__ == "__main__":
    l = [2,2,1,1,1,2,2]
    su = Solution()
    print(su.majorityElement1(l))