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



## 5、判断单链表是否有环？有环的话入环节点是什么？

快慢指针都从头开始，慢指针走一步，快指针走两步

​			--> 有环的话快慢指针一定会在环内相交, 相交后让慢指针回到head，然后慢指针和快指针都每次走一步，相交后即为入环节点

​          ---> 没环的话快指针会走向nil

## 6、LRU缓存

保证O(1)时间get, 使用map保存key，value为链表节点

保证能够在容量不足时候删除最久没有被使用的，就要有一个数据结构体来保存每个key访问的顺序，访问了就在调整到最前面，那么最后面的就是最长时间没有被访问的。所以使用双链表，若容量不够则删除聊吧尾部节点，然后在头部插入

为了保证删除最久没有被使用的，从链表尾删除，每次发生对节点的读取将该节点移动到链表头部



set(key)

key存在于map中，则从map的value拿出该链表节点，将其调整到头部

key不存在：

​           --> 容量未满：直接插入链表头，并放入map中

​           --> 容量满：   先删除链表尾部,再插入链表头，并放入map中



get(key)

key存在于map中, 从则从map的value拿出该链表节点，将其调整到头部, 然后返回value

key不存在，返回nil

## 7、20 有效的括号

（）{}[]， 输入字符串，输出是否是有效的括号

思想：利用成对出现的特性

法一：辅助栈，循环遍历字符串中字符，每来一个左括号入栈，可以是'('、'['或者’{‘

​                      每来一个右括号，

​									--> 栈空：return false

​    								--> 栈非空，栈顶pop一个并和当前右括号比较是否成对，不成对false

​          返回：循环结束后栈空则true, 非空则false

法二:  还是辅助栈，但是相对于法1左右括号是否成对的比较会更容易，来了一个左括号入栈的不是左括号，而是相应的右括号，这样当出栈之后只要比较字符是否相等即可。

## 8、32最长有效括号

dp

法一：一维dp, 这个题不能用二维dp，否则讨论起来会非常麻烦，必须利用该题的有效括号的特性来使用dp, 一个有效括号肯定得是以')'右括号结尾，

1、dp[i]定义为从头开始以s[i]为结尾的字符串的最长有效括号

2、递推公式：

```
 dp[i]   ---> 当s[i]  == '(' 为0
​        --->  s[i] == ')'  ---> s[i-1] == '(', 形如....(), 此时后面这两个括号一定是成对的，最大长度= 2+dp[i-2]
                            ---> s[i-1] == ')', 形如....)), 此时根本不知道这两个右括号和谁是成对的，利用dp[i-1]信息
                                 假设以s[i-1]结尾的最长有效括号长度为n且不为0，(n==0就不用讨论了,dp[i]=0)则与i-1位置成对出现的左括号的下标为i-n, 再看下标i-n-1处是有可能和i下标处成对出现的情形
                                            则. .  x     (   ...   )   )
                                                 i-n-1  i-n       i-1  i
                                 --->  s[i-n-1]  == '(', 此时形如....((...)) 则 dp[i] = n+2+dp[i-n-2], 其中n换成dp[i-1]
                                 --->  s[i-n-1]  == ')', 此时形如....)(...)) 则 没法配对dp[i] = 0
                                            
```

   时间O(N), 空间O(N)

其他方法比较难想

## 9、[3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters)

输入s, 输出int长度

法一：滑动窗口

1、定义窗口S[L:R]闭区间表示以s[i]为结尾的无重复子串，

2、循环更新：l和r两个指针初始为0， 让r从0开始递增一直到长度n-1结束，每次r右扩时候看当前从S[L:R]闭区间上的子串是否满足无重复性，若不满足则L向右前进缩小范围直到满足无重复性这个条件，R-L+1为当前以s[r]结尾的最长子串

3、怎么判断是否满足无重复性，可以用一个set来保存当前窗口上的值，若r的加入打破了无重复性说明s[r]已经在集合中存在了，那么让L每次前进一步直到S[r]不在集合里面了即可



时间O(N), 空间O(m） m最长字符串长度

code如下：

```go
func lengthOfLongestSubstring(s string) int {
	var l, res int
	set := make(map[byte]struct{}) // 当前窗口内的元素集合
	for r := 0; r < len(s); r++ {
		for { // 一直缩小窗口，直到当前的s[r]不在集合中
			_, ok := set[s[r]]
			if !ok {
				break
			}
			// set中删除当前l所在元素，然后l++, 窗口向右缩减
			delete(set, s[l])
			l++
		}
		set[s[r]] = struct{}{}
		// 此时的窗口满足以s[r]结尾的最长子串，进行结算
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
```

法二：对法一的窗口缩减进行加速

想一下在上面的窗口每次缩减的结果都是什么样的：

假设当前窗口为：abc，来了一个b此时abcb不满足无重复特性，l一直右移，直到窗口内字符串为cb,假设字符和下标对应如下：

```
字符：a b c b
下标：2 3,4 5
     L      R
```

发现没有L其实就是更新到b上一次出现的index+1处，所以L不用一步一步往右走，L可以直接跳到index+1

但是要考虑一种情况，当前的map定义不是窗口内的集合，而是0开始到r位置的集合，所在考虑下面这种情况

```
f....abcf
0    l  r
```

当前r来到f, f在集合里面，但是上一次的index=0, 此时我的l不用向右走的，维持l即可，所以会有下面第6行代码

```go
func lengthOfLongestSubstring(s string) int {
	var l, res int
	set := make(map[byte]int) // 此时这个map的定义不是窗口内的字符，而是从0开始到r为止出现的字符，有点空间换时间的意思
	for r := 0; r < len(s); r++ {
		if index, ok := set[s[r]]; ok {
			l = max(l, index+1) 
		}
		set[s[r]] = r
		// 此时的窗口满足以s[r]结尾的最长子串，进行结算
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
```

## 10、53最大连续子数组和

法一：dp

分析这道题要求的最大连续子数组，看到连续你就要往以什么结尾的方式来想，

**定义**：然后定义dp[i] 就是以第i个元素结尾的最大连续子数组和

**递推**：第i个元素的最大连续子数组和，和dp[i-1]是息息相关的，其实第i个元素只有两种选择：加入以i-1为结尾的连续子数组 或者 不加入

加入：nums[i] + dp[i-1]

不加入： nums[i]

那就看加入和不加入谁更大我就选谁，所以递推公式如下“

dp[i]  = max{ dp[i-1]+nums[i],  nums[i]}

**初始条件：**dp[0] = nums[0]



法二：分治， 和线段树有关，暂时没看



## 11、300最长递增子序列

法一：dp

定义：dp[i]为以nums[i]为结尾的最长递增子序列长度，长度为n, 答案为dp数组的最大值

递推：考察nums[i]， 以nums[i]结尾的最长长度，nums[i]肯定要加入，所以去i前面找比nums[i]小的元素中子序列最长的就是答案

```go
 dp[i]  = 1 + max{dp[j], j=i-1,i-2,....,0}
```

初始条件：dp[0] = 1



时间：O(N**2), 空间O(N)

​      







法二：贪心 + 二分查找，不好想

看看这个题解：https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326308/C%2B%2BPython-DP-Binary-Search-BIT-Segment-Tree-Solutions-Picture-explain-O(NlogN) 

step1: 定义sub1数组为最终的最长子序列，长度len动态扩展，刚开始为0，随着有数字加进来增长

step2: 遍历原数组nums, 

​								 ---->		若res数组为空，nums[i]进数组

​                            ---> 非空，此时考察当前这个nums[i]能否加入我们成为一员，使用贪心的方法来考察，

​												 ---> 若当前的nums[i]加入之后我们子序列长度变得比原来长了，欢迎加入

​                                         ---> 若加入之后长度和原来相等，但是把当前sub1的最大值变小了，同样欢迎加入，例如res=[2,3，5] 来了一个4， 可以变成234，使得序列的最大值由5变成4了，也欢迎，因为4比5更有可能在后面变得更长， 再例如res=[9], 来了一个1， 替换原来的9 res=[1]

​                                          ----> 若加入之后，虽然使得sub1长度变小了，但是我们不欢迎它加入sub1, 但是仍然要保留它，例如res=[2,3,5,8], nums[i]=4, 我们必须保留这个4，因为【2,3,4】这个序列的潜力更大，它以后可能会有2,3,4,5,6,7这种序列存储，所以此时我们保留2个数组，sub1=[2,3,5,8], sub2=[2,3,4]



这种考察是怎么进行的呢，由于sub是严格递增的，所以我们找nums[i]在sub中插入的位置index在哪，插入之后还是严格递增的，

​           若 index = len(sub), 说明nums[i]是最大的，使得序列变长了

​           若 index = len(sub)-1, 说明nums[i] < 当前sub的最大值，例如res=[2,3,5]  nums[i] = 4, 求得index=2, 此时替换res[index] = nums[i]

​           若 0 <= index < len(sub)-1, 增加一个新数组



上面的思路由于会增加多个sub1会导致性能很差，那么修改一下sub数组的定义就可以使得我们不用每次都新增sub数组，定义sub[k]是当前已经遍历的数组元素中递增子序列长度为k+1的所有子序列中尾部最小的那1个

例如nums = [982354]， 他的长度为1的递增子序列有2,3，45,8,9 都可以，取里面最小的2，所有sub[0] = 2

​                                      长度为2的递增子序列有[2,3], [2,5],[2,4], 取尾部最小的也就是[2,3]的尾部为3最小，所以sub[1]=3

​                                      长度为3的自增子序列有[2,35] 和234， 取4， 所以sub[2] =4

sub = [2,3,4]  sub的长度就是要求的长度





## 12、二分查找

要练代码，两种写法，[]和[), 保持两种原则，两种都要写

二分查到排序数组中，找到target在排序数组中应该存在的位置，可能是0，n, 
