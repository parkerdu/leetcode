# nums = list(input().split(","))
# # nums = ['32','2','56','9']
# def quick_sort(alist, first, last):
#     """快速排序"""
#     if first >= last:
#         return
#     mid_value = int(alist[first][0])  # 要定位的值
#     mid_str = alist[first]
#     low = first  # 左游标
#     high = last  # 右游标
#     while low < high:
#         #  第一步让high从右往左走
#         while high > low and int(alist[high][0]) >= mid_value:  # =号只放一个下面不放，实现和中间值相等的数都放在右边
#             high -= 1  # 退出循环就表示high走到了<=中间值的一个人面前
#         alist[low] = alist[high]
#         #  第二步high被复制后让low往右走
#         while low < high and int(alist[low][0]) < mid_value:
#             low += 1  # 退出循环表示low走到大于等于中间值那
#         alist[high] = alist[low]
#         # 第三部low被复制后转第一步
#     # 循环退出 low=high= 中间该在的位置
#     alist[low] = mid_str
#
#     quick_sort(alist, first, low-1)  # 对mid_value左边递归调用函数
#     quick_sort(alist, low+1, last)  # 对mid_value右边边递归调用函数
#
# quick_sort(nums,0,len(nums)-1)
# res = ''
# for i in nums[::-1]:
#     res += i
# print(res)

# nums = list(input().split(","))
nums = ['30','3']
a = sorted(nums,reverse=True)
# print(a)
res = ''
for i in a:
    res += i

print(res)
