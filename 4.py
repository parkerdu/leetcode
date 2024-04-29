class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        """
        我的解法：时间O(M+N)
        把两个数字用merge_sort的合并方法合起来，用时O(M+N),从结果上可以计算时间复杂度，最终的
        merge_list 长度是M+N,你每次往里面append一个数
        :param nums1:
        :param nums2:
        :return:
        """
        m = len(nums1)
        n = len(nums2)
        if (m+n) & 1:
            index = [int((m+n)/2)]
        else:
            index = [int((m+n)/2-1),int((m+n)/2)]
        length = index[-1] + 1
        left = 0
        right = 0
        merge_list = []
        while left < m and right < n:
            if len(merge_list) == length:
                break
            if nums1[left] <= nums2[right]:
                merge_list.append(nums1[left])
                left += 1
            elif nums1[left] > nums2[right]:
                merge_list.append(nums2[right])
                right += 1
        merge_list += nums1[left:]
        merge_list += nums2[right:]
        if len(index) == 1:
            return merge_list[index[0]]
        else:
            return (merge_list[index[0]] + merge_list[index[1]])/2

    def findMedianSortedArrays1(self, x, y):
        """
        最优解法：时间O(log(min(M,N))),M, N = len(x), len(y)
        具体视频讲解：https://www.youtube.com/watch?time_continue=1428&v=LPFhl65R7ww
        做法：把两个数组各自拆分成两个子数组
        x = x_left + x_right\
        y = y_left + y_right
        使得left中的所有数都小于right中的数，且如果m+n是偶数，len(x_left+y_left) = len(x_right+y_right)
                                            如果m+n 是奇书，len(x_left+y_left) = len(x_right+y_right)+1
        如果找到上面这样的划分，则中位数= 如果m+n是偶数，max(x_left, y_left)+min(x_right，y_right) /2
                                            如果m+n 是奇书，max(x_left, y_left)
        :param self:
        :param nums1:
        :param nums2:
        :return:
        """
        m, n = len(x), len(y)
        if n < m:
            return self.findMedianSortedArrays1(y,x)
        first = 0
        last = m
        while first <= last:
            p_x = (first + last)//2

            x_left, x_right = x[:p_x], x[p_x:]
            p_y = (m+n+1)//2 - p_x
            y_left, y_right = y[:p_y], y[p_y:]
            if not x_left:
                x_left = [float("-inf")]
            if not x_right:
                x_right = [float("inf")]
            if not y_left:
                y_left = [float("-inf")]
            if not y_right:
                y_right = [float("inf")]
            if y_left[-1] > x_right[0]:
                # 往右边走
                first = p_x +1
            elif x_left[-1] > y_right[0]:
                # 往左边走
                last = p_x -1
            elif x_left[-1] <= y_right[0] and y_left[-1] <= x_right[0]:
                if (m+n)%2 == 0:
                    return (max(x_left[-1],y_left[-1])+min(x_right[0],y_right[0]))/2
                else:
                    return max(x_left[-1],y_left[-1])




if __name__ == "__main__":
    a = [1]
    b = [2,3]
    su = Solution()
    print(su.findMedianSortedArrays1(a,b))
    print(max(float("-inf"),56))