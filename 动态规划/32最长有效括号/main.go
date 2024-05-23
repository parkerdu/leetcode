package main

import "fmt"

//func longestValidParentheses(s string) int {
//	// dp[i][j]为s[i:j]是不是有效的，n*n的数组
//	n := len(s)
//	dp := make([][]int, n)
//	for i := range dp {
//		dp[i] = make([]int, n)
//	}
//	max := 0
//	for i := n - 2; i >= 0; i-- {
//		for j := i; j < n; j++ {
//			if s[i] == '(' && s[j] == ')' {
//				if i+1 == j {
//					dp[i][j] = 1
//				} else {
//					if s[i+1] == '(' { // ((.....)
//						dp[i][j] = dp[i+1][j-1]
//						if j-2 > i {
//							dp[i][j] += dp[i][j-2]
//						}
//					} else if i+2 < j { // ()....)
//						dp[i][j] = dp[i+2][j]
//					}
//				}
//				if dp[i][j] == 1 && j-i+1 > max {
//					max = j - i + 1
//				}
//			}
//		}
//	}
//	return max
//}

/*
dp[i]   ---> 当s[i]  == '(' 为0
​        --->  s[i] == ')'  ---> s[i-1] == '(', 形如....(), 此时后面这两个括号一定是成对的，最大长度= 2+dp[i-2]

	---> s[i-1] == ')', 形如....)), 此时根本不知道这两个右括号和谁是成对的，利用dp[i-1]信息
	     假设以s[i-1]结尾的最长有效括号长度为n,则与i-1位置成对出现的左括号的下标为i-n, 再看i-n-1
	                则. .  x     (   ...   )   )
	                     i-n-1  i-n       i-1  i
	     --->  s[i-n-1]  == '(', 此时形如....((...)) 则 dp[i] = n+2+dp[i-n-2], 其中n换成dp[i-1]
	     --->  s[i-n-1]  == ')', 此时形如....)(...)) 则 dp[i] = 0
*/
func longestValidParentheses(s string) int {
	// 定义dp[i]为以i为结尾的字符串的最长有效括号
	n := len(s)
	dp := make([]int, n)
	max := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // ....()
				dp[i] += 2
				if i-2 > 0 {
					dp[i] += dp[i-2]
				}
			} else { // ....))
				n := dp[i-1]
				if n != 0 && i-n-1 >= 0 && s[i-n-1] == '(' {
					dp[i] += n + 2
					if i-n-2 > 0 {
						dp[i] += dp[i-n-2]
					}
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max

}

func main() {
	s := "((()))())"
	fmt.Println(longestValidParentheses(s))
}
