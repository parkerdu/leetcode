# n = int(input())
# condition = list(map(int,input().split()))
n = 1000
condition = []
for i in range(999):
    condition.append(1)
nums = []
for i in range(1,n+1):
    nums.append(i)

class Aiqiyi():

    def permute(self):
        self.res = 0
        self.dfs(nums, [])
        print(self.res%(10**9+7))

    def dfs(self, nums, path):
        if not nums:
            self.res += 1
        index = len(path) - 1
        for i in range(len(nums)):
            if index >= 0:
                if condition[index] == 1 and path[index] > nums[i]:
                    self.dfs(nums[:i] + nums[i + 1:], path + [nums[i]])
                elif condition[index] == 0 and path[index] < nums[i]:
                    self.dfs(nums[:i] + nums[i + 1:], path + [nums[i]])

            else:
                self.dfs(nums[:i] + nums[i + 1:], path + [nums[i]])


if __name__ == "__main__":
    su = Aiqiyi()
    su.permute()