class Solution:
    def threeSum(self, nums):
        dict_map = {}
        results = []
        for i,x in enumerate(nums):
            for j in range(i+1,len(nums)):
                """这种做法出现一个问题，就是你dict_map中存的是c=-a-b的值，有可能对同一个c值由两组不同的a,b你只能存一组"""
                target = -nums[j]
                # x在匹配中，并且下标不同
                if x in dict_map and i not in dict_map.get(x) and len(dict_map.get(x)) == 2:
                    index = dict_map.get(x)
                    index.append(i)
                    result = []
                    for i in index:
                        result.append(nums[i])
                    result = sorted(result)
                    if result not in results:
                        results.append(result)

                else:
                    dict_map[target-x] = [i,j]

        return results
    def threeSum1(self, nums):
        """法二：转化成两数之和来做
            两数之和再nums中找a+b=target
            三数之和在nums中找a+b+c=0
                转化为b+c=-a=target 也就是说目标现在是动的
                那么我们先遍历整个列表，把第0个数作为目标，在剩下的列表（除去第0个数）中就是解决
                    b+c = target的两数之和问题
            时间复杂度：遍历整个列表O(N)
                        在剩下的列表中解决两数之和O(N)
                    总时间O(N**2)
            空间复杂度：O(N)"""
        """tansform the question from threesum to twosum
            twosum:a+b=target
            threesum:a+b+c=0
                b+c=-a=target compare to the twosum problem this target in threesum is not fixed
                so we can traval the whole list,
                step1:treat the nums[i] as minus target
                step2:the sublist nums[i+1:] with a target which is fixed now in step1 is a twosum problem
            time complexity:O(N^2)
            SPACE complexity:O(N)"""
        nums.sort()
        results = set()
        for i,a in enumerate(nums[:-2]):
            target = -a
            dict_map = {}
            # now the target is fixed,we solve the twosum problem
            # 把目标固定，变成在剩下的列表中解决两数之和问题
            for j,b in enumerate(nums[i+1:]):
                if b in dict_map:
                    results.add((a,b,target-b))
                else:
                    dict_map[target - b] = j
        # map,把集合转化为列表
        return list(map(list,results))

    def threeSum2(self, nums):
        """法二的优化，同样是转化成两数之和来做，加两行代码执行的时间就大大减少
            因为，你先对列表进行排序了，所以如果存在两个数相等，那后面那个数就不用在做了
            例如a = [-2,-2, 0, 1,1,1,1,1]
            你先把目标设置为2，在剩下列表[-2, 0, 1,1,1,1,1]中找两数之和为2
              再把目标设置为2，在剩下列表[0, 1,1,1,1,1]中找两数之和为2
              很明显可以看到第二次找出来的结果一定是第一次的子集，所以根本不用再去做第二步操作了，从而大大减少运行时间
              当然从时间复杂度上来说两种方法都是O(N**2)的，但是优化之后的最优时间是减少的"""
        nums.sort()
        results = set()
        for i,a in enumerate(nums[:-2]):
            if i >= 1 and a == nums[i-1]:
                continue
            target = -a
            dict_map = {}
            # 把目标固定，变成在剩下的列表中解决两数之和问题
            for j,b in enumerate(nums[i+1:]):
                if b in dict_map:
                    results.add((a,b,target-b))
                else:
                    dict_map[target - b] = j

        return list(map(list,results))

    def threeSum3(self, nums):
        """老师的法二"""
        if len(nums) < 3:
            return []

        nums.sort()
        res = set()
        for i,v in enumerate(nums[:-2]):
            if i >= 1 and v == nums[i-1]:
                continue
            d = {}
            for x in nums[i+1:]:
                if x not in d:
                    d[-v-x] = 1
                else:
                    res.add((v,-v-x,x))

        return list(map(list,res))

if __name__ == "__main__":
    nums = [2,-1,-1,-1,-1,-1,-1,0]
    su = Solution()
    print(su.threeSum1(nums))
    print([-1, 1, 0]==[-1, 1,0])