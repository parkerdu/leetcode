class Solution:
    def maxSlidingWindow(self, nums, k):
        """滑动窗口中的最大值
            思想：双端队列
            windows 用来存储列表中的下标，
                满足下面条件：1.windows长度最大为3
                             2.windows中下标对应的数是从大到小排的，例如nums[windows[0]]>nums[windows[1]]>nums[windows[2]]
                             这是因为每次新进来一个人，都要从右往左（pop)杀掉所有比他小的的人
                             3.新进来下标从后面append
            res 用来存储当前窗口中的最大值
            总结：每次窗口向右走一步，但是window不把所有的三个下标都存下来
                    a.先删除老下标
                    b.新人下标一定存，但是存之间先把以前那些不如我的全部赶走（赶走并不影响我找最大值），
                    所以windows里面经常都是小于三个数
                             """
        if not nums: return []
        windows, res = [], []
        for i,x in enumerate(nums):
            #  i>=k表示只有下标超过k时候才需要移动
            #  i-k+1是当前窗口的最小下标，所有比他还小的都要丢掉了
            #  <i-k+1 等价于 <= i-k
            if i>=k and windows[0]<= i-k:
                windows.pop(0)
            while windows and x>nums[windows[-1]]:
                windows.pop()
            # 每次我肯定是都会把下标i添加到windows中去，但是在添加之前的两个操作，
            # 一个是超过的窗口的老元素去掉，一个是前朝不如我的老臣杀掉
            # 所以最终保证了上面绿色的三个条件
            windows.append(i)
            # i=2,3,4,5,....，n-1都要添加一个最大值
            if i >= k-1:
                res.append(nums[windows[0]])
        return res