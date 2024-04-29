class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        self.results = []
        self.min = min(candidates)
        if target in candidates:
            self.results.append([target])

        self.help(candidates,target,[])
        return list(self.results)

    def help(self,candidates, target, result):
        if target < 0:

            return
        if target == 0:
            result = sorted(result)
            if result not in self.results:
                self.results.append(result)
        for i in candidates:
            self.help(candidates,target-i,result+[i])

class Solution1(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        self.results = []
        self.help(sorted(candidates),target,[])
        return list(self.results)

    def help(self,candidates, target, result):
        if target < 0:

            return
        if target == 0:
            result = sorted(result)
            if result not in self.results:
                self.results.append(result)
        for i in candidates:
            self.help(candidates,target-i,result+[i])
if __name__ == "__main__":
    a = [3,2]
    su = Solution1()
    print(su.combinationSum(a,6))