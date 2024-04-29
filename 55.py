#coding:utf-8
class Solution:
    """55题跳跃游戏"""
    def canJump(self, nums) -> bool:
        """我的解法，遇到0判断能否跳过去"""
        for i in range(len(nums)-1):
            if nums[i] > 0:
                i += 1
            else:
                j = i - 1
                count = 0
                while j >= 0:
                    if j+nums[j] > i:
                        break
                    else:
                        count += 1
                    j -= 1
                # 如果可以过去的话循环结束count<i
                # 如果等于i说明没有一个人可以越过0
                if count == i:
                    return False
        return True

    def canJump1(self, nums):
        """法二：每次记录当前可以跳的最大index在哪
        然后判断当前点时候可以过去，如果当前点都不能过去，那么一定去不了最后"""
        # m记录当前可以到达的最大index
        m = 0
        for i, n in enumerate(nums):
            if i > m:
                return False
            m = max(m, i + n)
        return True

if __name__ == "__main__":
    a = [2,0,0]
    su = Solution()
    print(su.canJump2(a))