# Leetcode-算法与数据结构

该仓库记录了从2018年至今的刷题记录，包含自己刚开始无头苍蝇一样的随机刷题，到后来跟着覃超系列刷题，再后面工作了跟着左神的视频进行系统的学习，慢慢开始把原来的点开始构建成线，但到2024年找工作始终位还是未构成完成的面，算法的学习是一个持续的事情，要循序渐进，始终学习。



2024面试拼多多时候发现，是一个曾写过的单词搜索问题，但是在面试的场景下，还是想不起来，所以制定了对旧题的复习计划，每天3个题目，2个新题，1个旧题，并总结每个题目的关键词。

# 1、2024年刷题记录

## 2024.04.29

**2个新题：**42接雨水 +128 数组最长连续元素

**1个旧题：**堆排序

**关键词：**

接雨水核心：拆分成子问题，每一根柱子上能装多少水，加起来就是总的水量

数组最长连续元素：借助哈希表的O(1)查询能力，先找到连续数组的起点，然后再找能持续多长

堆排序：堆节点index结构，向上弑君，向下除子

## 2024.04.30

**2个新题：**92反转链表  + 199二叉树右视图

**1个旧题：**小和问题

**关键词：**

反转链表：定义好4个关键变量

二叉树右视图：层次遍历 或者 根右左递归

小和问题：和接雨水有点类似，要看子问题，每个数为小和贡献了多少次，例如 1,3,2,5 ,4
对于1来说，右边有3个比他大 sum += 1*3
对于3来说 sum+= 3*1
对于2， sum+= 2*2
这样想之后才能将其和归并排序结合起来。

## 2024.05.03

**2个新题:** 297二叉树序列化和反序列化 + 200岛屿个数

**1个旧题：**493 逆序对问题

**关键词：**

二叉树序列化和反序列化：孰能生巧 + 字符串下标直接可以找到left节点

岛屿个数： 利用二维数组降低为一维的技巧 利用并查集

逆序对：转换为大和问题



## 2024.05.09

1个旧题: 逆序对问题



## 2024.05.11

1个旧题：25k个一组反转链表 (该题今天面试考到了，现场虽然写出来，但是不够熟练)

关键点：

1、4个变量，之前的头和尾巴，反转后的头，反转后的start, stop，主函数里面两组之间拼接，rt.next, rt = start,stop, 不要忘了rt自己也要更新

2，子函数要返回3个参数，start,stop,next 前两个给组之间合并用，next表示当前head链表的未遍历的第一个元素，给head更新用

## 2024.05.13

92链表翻转2 + k个一组翻转  -- 已完成

https://github.com/lewiscrow/WorkHardAndFindJob/blob/master/%E5%A4%8D%E4%B9%A0/%E9%9D%A2%E8%AF%95/%E5%AD%97%E8%8A%82%E8%B7%B3%E5%8A%A8%E9%9D%A2%E7%BB%8F.md



| 215. 数组中的第K个最大元素 | 3                                                         |
| -------------------------- | --------------------------------------------------------- |
| 145. 二叉树的后序遍历      |                                                           |
| 33. 搜索旋转排序数组       | 按mid在左半边还是右半边来分，在左半边再讨论，右半边再讨论 |



33题关键点：

![](images/33.jpg)

- If the entire left part is monotonically increasing, which means the pivot point is on the right part
  - If left <= target < mid ------> drop the right half
  - Else ------> drop the left half
- If the entire right part is monotonically increasing, which means the pivot point is on the left part
  - If mid < target <= right ------> drop the left half
  - Else ------> drop the right half

参考链接：https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14436/Revised-Binary-Search/191339

```java
public int search(int[] nums, int target) {
    if (nums == null || nums.length == 0) {
        return -1;
    }
    
    /*.*/
    int left = 0, right = nums.length - 1;
    //when we use the condition "left <= right", we do not need to determine if nums[left] == target
    //in outside of loop, because the jumping condition is left > right, we will have the determination
    //condition if(target == nums[mid]) inside of loop
    while (left <= right) {
        //left bias
        int mid = left + (right - left) / 2;
        if (target == nums[mid]) {
            return mid;
        }
        //if left part is monotonically increasing, or the pivot point is on the right part
        if (nums[left] <= nums[mid]) {
            //must use "<=" at here since we need to make sure target is in the left part,
            //then safely drop the right part
            if (nums[left] <= target && target < nums[mid]) {
                right = mid - 1;
            }
            else {
                //right bias
                left = mid + 1;
            }
        }

        //if right part is monotonically increasing, or the pivot point is on the left part
        else {
            //must use "<=" at here since we need to make sure target is in the right part,
            //then safely drop the left part
            if (nums[mid] < target && target <= nums[right]) {
                left = mid + 1;
            }
            else {
                right = mid - 1;
            }
        }
    }
    return -1;
}
```









## 2024.05.14

- LRU


- 72编辑距离

关键点: 用临时变量保存dp[j], 更新时候用leftup, 更新完成后给leftup赋值 + 替换，插入，删除都要考虑

- 557反转字符串中的单词 III



- 2最长有效括号

快排

环形链表





# todo 旧题： 

给一个整数数组，找右边元素减左边元素的最大差值

旋转数组

32最长有效括号

- 88 合并两个有序数组



# todo新题

打家劫舍123

2个新题：洪水填充+ [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

# 阿里后端题目

[参考leetcodeTOP](https://github.com/afatcoder/LeetcodeTop/blob/master/alibaba/backend.md)

网站题目频率排序 https://codetop.cc/home

| 算法题                                  | 次数 |
| --------------------------------------- | ---- |
| 215. 数组中的第K个最大元素              | 3    |
| 145. 二叉树的后序遍历                   | 2    |
| 1. 两数之和                             | 2    |
| 232. 用栈实现队列                       | 2    |
| 88. 合并两个有序数组                    | 2    |
| 15. 三数之和                            | 2    |
| 349. 两个数组的交集                     | 1    |
| 5. 最长回文子串                         | 1    |
| 557. 反转字符串中的单词 III             | 1    |
| 72. 编辑距离                            | 1    |
| 543. 二叉树的直径                       | 1    |
| 144. 二叉树的前序遍历                   | 1    |
| 94. 二叉树的中序遍历                    | 1    |
| 剑指 Offer 22. 链表中倒数第k个节点      | 1    |
| 2. 两数相加                             | 1    |
| 70. 爬楼梯                              | 1    |
| 509. 斐波那契数                         | 1    |
| 21. 合并两个有序链表                    | 1    |
| 46. 全排列                              | 1    |
| 82. 删除排序链表中的重复元素 II         | 1    |
| 剑指 Offer 52. 两个链表的第一个公共节点 | 1    |
| 225. 用队列实现栈                       | 1    |
| 350. 两个数组的交集 II                  | 1    |
| 415. 字符串相加                         | 1    |
| 268. 缺失数字                           | 1    |
| 53. 最大子序和                          | 1    |
| 344. 反转字符串                         | 1    |
| 136. 只出现一次的数字                   | 1    |
| 剑指 Offer 24. 反转链表                 | 1    |

# 题目思路

## 1、求完全二叉树的最后一个节点

法一：层次遍历，时间O(N)

法二：利用完全二叉树index的特点，最后一个节点的index=n-1, 从最后一个节点开始向上找父亲节点的索引，一直找到根节点为止，这样我就有了从最后一个节点到root的索引信息，然后利用完全二叉树左节点的索引为奇数，右节点的索引为偶数的性质，就可以将得到的索引，转换为从左边还是走右边的路 

## 2、15三数之和为0

法一：将整个数组分成三部分，p正数，z零，n负数

总共有4种组合可能为0：正+正+负、正+零+负、正+负+负、零+零+零

所以将正数和负数变成set, 然后按照上述4种组合去求，时间复杂度变成O(N**2), 

正+正+负 ： 遍历正数的组合，去负数集合里面查找相反数在不在O(N**2)

正+零+负：若有0，则遍历正数的元素，有相应的负数即可配对，O(N)

正+负+负、

零+零+零

时间O(N**2), 空间O(N)

法二：时间O(N**2), 空间O(1)

step1: 排序原数组

step2: 遍历排序好的数组，若nums[i] = nums[i-1], 表示之前已经找过一样的数，跳过

step3: 使用l,r两个指针，l=i+1, r=n-1, n为长度，l<r则一直循环step4

step4: i+l+r < 0, 代表l+r不够大，l++，l往大数方向走， 

> 0 r--，  ==0， 说明当前组合符合条件，此时为了避免重复，l和l+1相等的话l还要往右边走，同理r,最后l++,r--

## 3、349两个有序数组的交集

法1： 时间O(M+N)

两个指针：i,j  i<j  i++

​                   i==j 记录并且i++,j++

​                  i>j  j++

知道有一个数组越界为止



法1在有一个数组很长，一个数组很短时候时间负责度不是很好



法二：遍历短的数组元素，去长的数组里面做二分查找，查到说明在

时间O(m*logN)

## 4、5最长回文子串

法一：时间O(N**2), 空间O(1)

遍历字符串字符，使用两个指针l,r， l=i, r=i 和 l=i,r=i+1(保证了aa这种情况也可以顺利外扩)都去做一遍外扩动作，长度为r-l+1

法二：dp

dp[i] [j]定位为s[i:j] 闭区间上是不是回文子串，最终返回dp数组中所有为true中j-i+1最大者

递推公式： s[i ]== s[j]  ---> j=i+1, 是的

​                               ----> 否则若还满足dp[i+1] [j-1]也是，则dp[i] [j]为true

初始条件： 所有对角线上均为true,

更新方向：从底向上更新



法三：manacher算法

p数组，p[i] = 以i为中心的半径

r当前的回文最右边界，以整个数组维度的

c是r对应的最早中心点



遍历数组，i>=r, 则从i开始外扩

i<r  --> 若i'的回文半径在r范围之内：p[i] = p[i']

​     --> 若i'的回文半径超过r返回，r-i,从i开始外扩只能扩到r, 超出了说明当前r算的不对，r还能扩

​     --> 刚好和r重合，从r开始外扩



coding: 每次从i开始外扩时候，给一个初始值 i>=r len=1

​                                                             --> i < r len = math.min(r-i, p[i'])

从len开始向两边扩 

