import heapq
class KthLargest:
    """第k个最大值
        思想：把前k个最大数取出来，变成一个小顶堆，我们一直维护这个小顶堆
        小顶堆的最小值就是我们要的返回值，第k个最大值
        """

    def __init__(self, k: int, nums):
        """取出前k个最大数，法一：可以先排序再取前k个值，用时O(nlogn)
                        法二：用heapq.nlargest函数，但是其内部跟法一的做法是等价的，也是用sorted函数，所以时间复杂度也是O(NlonN)"""
        # 法一
        self.k = k
        self.nums = nums
        # python内置sorted函数会生成一个新列表，空间复杂度O(N)
        a = sorted(self.nums,reverse=True)
        self.heap = a[:self.k]
        heapq.heapify(self.heap)

        # 法二
        # self.k = k
        # self.nums = nums
        # self.heap = heapq.nlargest(self.k, self.nums)
        # heapq.heapify(self.heap)


    def add(self, val: int) -> int:
        # 如果刚开始不够k个数，就一直往里面push不用返回,直到有k个数了，开始返回
        if len(self.heap) < self.k:
            heapq.heappush(self.heap,val)
            if len(self.heap) == self.k:
                return self.heap[0]
        else:
            min = self.heap[0]
            if val > min:
                # heapq.heappop(self.heap)
                heapq.heapreplace(self.heap,val)
                return self.heap[0]
            else:
                return min

class KthLargest1:
    """优化版本，
    1.讨论情况根本不需要那么复杂
        只有两种情况需要改变堆，其他情况直接返回堆的第0个值就行了
        情况1：堆元素少于k个，一直加
        情况2：来的元素大于当前堆的最小值，把最小值拿走，加入来的人
        其他情况都不用动，然后直接返回改过的最小值就行了
    2.取前k个最大值还可以优化，
        heapq.heappop时间复杂度是O(logn)
        需要进行n-k次所以总时间复杂度小于O(NlogN)
        """

    def __init__(self, k: int, nums):
        self.k = k
        self.heap = nums
        heapq.heapify(self.heap)
        while len(self.heap) > self.k:
            heapq.heappop(self.heap)

    def add(self, val: int) -> int:
        # 如果刚开始不够k个数，就一直往里面push不用返回,直到有k个数了，开始返回
        if len(self.heap) < self.k:
            heapq.heappush(self.heap,val)
        elif val > self.heap[0]:
            heapq.heapreplace(self.heap,val)
        return self.heap[0]

# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)
if __name__ == "__main__":
    k = 3
    arr = []
    obj = KthLargest(k, arr)

    print(obj.add(3))
    print(obj.add(5))
    print(obj.add(10))
    print(obj.add(1))