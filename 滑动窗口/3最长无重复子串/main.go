package main

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

func lengthOfLongestSubstring(s string) int {
	var l, res, r int
	set := make(map[byte]struct{}) // 当前窗口内的元素集合
	for r < len(s) {
		_, ok := set[s[r]]
		if ok {
			// 在里面不满足条件
			delete(set, s[l])
			l++
		} else {
			// 不在里面结算
			set[s[r]] = struct{}{}
			// 此时的窗口满足以s[r]结尾的最长子串，进行结算
			if r-l+1 > res {
				res = r - l + 1
			}
			r++
		}

	}
	return res
}

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
