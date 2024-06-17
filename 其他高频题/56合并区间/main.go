package main

import "sort"

/*
此题的关键是排完序之后呈现的特性：能合并的区间一定是连在一起的，第0个区间能合并的一定是第1个区间，如果还能再延长就是第2个区间
因为此时min是按顺序增长的，所以第0个区间的min就决定了下限不会再往左走，能变的只有上限
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		// min1  min2  max1
		if res[len(res)-1][1] >= intervals[i][0] {
			// 可合并，上界更新为两个max的最大值
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
