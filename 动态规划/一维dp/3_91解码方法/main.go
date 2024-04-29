package main

import (
	"fmt"
	"runtime"
)

func numDecodings(s string) int {
	// step2: 定义dp数组，其中dp[i]定义为s[i:]的最大解码数，长度为len(n)+1
	n := len(s)
	dp := make([]int, n+1)
	// step3: 初始值dp[n] = 1
	dp[n] = 1
	// step4: 从右向左遍历s, 利用状态转移函数求解，
	for i := n - 1; i >= 0 ; i-- {
		if s[i] == '0' {
			continue
		}
		dp[i] = dp[i+1]
		fmt.Println(int(s[i]))
		if (i+1 < n) && (1 <= 10*int(s[i]-'0')+int(s[i+1]-'0')) && (10*int(s[i]-'0')+int(s[i+1]-'0') <= 26) {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(numDecodings("10"))
}

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 1<<16; i++ {
			fmt.Println(i)
		}
	}()
	for {}
}