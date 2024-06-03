package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	curLevelNode, nextLevelNode := 1, 0
	line := make([]int, 0)
	var reverse bool
	for len(queue) != 0 {
		// step1: 出队当前cur，curLevelNode--, 加入到line，reverse从右向左添加
		cur := queue[0]
		queue = queue[1:]
		curLevelNode--
		line = append(line, cur.Val)
		// step2: 左非空入队，右非空入队，同步更新nextLevelNode
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextLevelNode++
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextLevelNode++
		}
		// step3: 判断是否要结算,是转step4, 否则继续
		if curLevelNode == 0 {
			if reverse {
				// 反转当前层的数组
				for i := 0; i < len(line)/2; i++ {
					line[i], line[len(line)-1-i] = line[len(line)-1-i], line[i]
				}
			}
			reverse = !reverse
			res = append(res, line)
			curLevelNode, nextLevelNode = nextLevelNode, 0
			line = make([]int, 0)
		}
		// step4: 当前line加入res, reverse取反，更新curLevelNode，重置line和nextLevelnode
	}
	return res
}
