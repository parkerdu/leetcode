
class Solution():

    def find_max_list(self,nums):
        """链接：https://www.nowcoder.com/questionTerminal/8dd82ad329f249658a95b339b44ef887
            来源：牛客网
            给出一个数列（元素个数不多于100），数列元素均为负整数、0、正整数。请找出数列中的一个连续子数列，
            使得这个子数列中包含的所有元素之和最大，在和最大的前提下还要求该子数列包含的元素个数最多，
            并输出这个最大和以及该连续子数列的个数。

            解法：
            res = 0（res用来存储最大的和），只有当temp>res时候才会更新res
            遍历整个数组，用temp存储当前子序列的和，左右两个指针表示子序列的下标
                一个关键点：一旦遍历到第j个数使得当前子序列的temp<0,那么res不用更新还保持上次的res（这个res如果后面没有子序列大于他
                    就是最大的和，所以不可更新他）

                    更新i = j +1,
                    更新temp = 0,开始从第j+1个数重新计算temp
                两个条件：
                    1. 遍历到第j个元素，当前子序列和大于res,更新res,更新左右两个下标指针
                    2. 遍历到第j个元素，当前子序列和等于res并且 j-i > r - j(当前子序列的和上一步的和一样但是，比上一个子序列长）
                    只用更新两个下标指针即可
                """
        # 定义四个变量temp = sum([i,j]) res = sum([l,r]) 循环时候用i更新l,用j更新r
        temp, res, l, r = 0, 0, 0, 0
        i = 0
        for j in range(len(nums)):
            temp = temp + nums[j]
            if temp > res:
                # 只有temp > res这件事发生，才会去更新res
                res = temp
                l = i
                r = j
            elif temp == res and j-i > r - j:
                l = i
                r = j
            else:
                # 如果 temp < res,这时候没必要更新l,r,但是一旦temp < 0
                # 代表当前[i,j]这个子数组的和小于0，那么这一段数组我直接不要了，更新i,temp
                if temp < 0:
                    i = j + 1
                    temp = 0
        return res,l,r


if __name__ == "__main__":
    a = [6, -7, -3, -2, 4]
    su = Solution()
    print(su.find_max_list(a))