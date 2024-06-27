package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	res := make([]string, 0)
	process(s, 4, "", &res)
	return res
}

func process(s string, l int, cur string, res *[]string) {
	if l == 0 {
		*res = append(*res, cur[:len(cur)-1])
		return
	}
	// 判断当前的l是否满足
	n := len(s)
	start := max(1, n-3*(l-1))
	for i := start; i <= 3 && i <= n; i++ {
		if (i > 1 && s[0] == '0') || i == 0 {
			break
		}
		head := s[:i]
		val, err := strconv.ParseInt(head, 10, 64)
		if err != nil {
			break
		}
		if 0 <= val && val <= 255 {
			next := cur + head + "."
			process(s[i:], l-1, next, res)
		}
		// 满足加上拼接前面的
		// 不满足返回

	}
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
