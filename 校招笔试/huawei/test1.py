#coding=utf-8
# 本题为考试单行多行输入输出规范示例，无需提交，不计分。
# 读取单行数据
# input：2 3
# output: 5
# import sys
# for line in sys.stdin:
#     a = line.split()
#     print(int(a[0]) + int(a[1]))
#
#
# def get_sum():
#     try:
#         while True:
#             a = list(map(int, input().split(" ")))
#             print(a[0]+a[1])
#     except:
#         pass
# get_sum()
a = '001'
print(int(a[::-1]))

def reverseAdd(a,b):
    if a < 1 or a > 70000:
        a = -1
    if b < 1 or b > 70000:
        b = -1
    a = str(a)
    b = str(b)
    return int(a[::-1]) + int(b[::-1])

a = 123
b = 456
print(reverseAdd(a,b))