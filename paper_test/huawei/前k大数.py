import heapq


class Solution(object):
    def findKthLargest(self, nums, k):
        """
        我的做法：大顶堆
        数组变成大顶堆，删除前k-1个堆顶元素，当前堆得堆顶取相反数即为第k大数
        :type nums: List[int]
        :type k: int
        :rtype: int
        heapq.heappop时间复杂度是O(logn)
        变成相反数用了O(N)
        生成小顶堆用O(N) heapq.heapify(nums)
        删除前k-1个最大数，每次用时O(logN) 删了k-1次，
        总时间：2*N+(K-1)*log(N)
        """
        # 变相反数，O(N)
        for i in range(len(nums)):
            nums[i] = -nums[i]
        # 生成小顶堆：O(N)
        heapq.heapify(nums)
        # 删除k-1次，每次用时 O(logN)
        for _ in range(k-1):
            heapq.heappop(nums)
        return -heapq.heappop(nums)

    def qucik_sort(self,nums, first,last,k):
        """
        快排思想做法，有人说是 O(N)时间，
        快排每次会把数组分成两部分nums = [left] +[mid_value] +[right]
        如果第k大数的index > nums.index(mid_value) 那就再取右边的数组中找，左边数组不用管了
        第一次遍历数组 用来O(N)时间
        后面每一次用的时间都会小于O(N)
        最多需要遍历log(N)次,但是只有第一次需要O(N)时间，所以总时间接近O(N)

        :param nums:
        :param first:
        :param last:
        :param k:
        :return:
        """
        if first >= last:
            return nums[first]
        mid_value = nums[first]
        left = first
        right = last

        while left < right:
            while left < right and nums[right] >= mid_value:
                right -= 1
            # 发生交换，右边拉到前面了
            nums[left] = nums[right]

            # 一旦发生交换，判断左指针
            while left < right and nums[left] < mid_value:
                left += 1
            nums[right] = nums[left]
        nums[left] = mid_value
        if k < left:
            return self.qucik_sort(nums,first,left-1,k)
        elif k == left:

            return nums[left]

        else:
            return self.qucik_sort(nums,left+1, last,k)



if __name__ == "__main__":
    a = [3, 2, 1, 5, 6, 4]
    su = Solution()
    print(su.qucik_sort(a, 0, len(a)-1, len(a)-3))
    print(a)
