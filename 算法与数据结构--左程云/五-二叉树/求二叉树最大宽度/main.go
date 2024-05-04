package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// step1: 准备3个变量 + 1个map + 1个队列, curLevel代表当前层默认为1， curLevelNode为当前层个数，默认为0，弹出才加1， max
	// map里面默认保存根节点和根节点层数
	curLevel, curLevelNode, max := 1, 0, -1
	levelMap := make(map[*TreeNode]int, 16)
	queue := make([]*TreeNode, 0, 16)

	// step2: 根节点入队列queue
	queue = append(queue, root)
	levelMap[root] = curLevel
	// step3: 若queue非空，弹出一个cur, 若cur所在层数为当前层，curLevelNode加1， 否则转step4
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		level := levelMap[cur]
		if level == curLevel {
			curLevelNode ++
		} else {
			// step4: 到这里说明进入下一层，进行结算，更新max, curLevel +1, curLevelNode = 1
			if curLevelNode > max {
				max = curLevelNode
			}
			curLevel++
			curLevelNode = 1
		}
		// step5: cur左孩子进队，且记录左孩子的层数，右孩子后进
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			levelMap[cur.Left] = level+1
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			levelMap[cur.Right] = level + 1
		}
	}


	// step6: 取max 和 curLevelNode最大值返回
	if curLevelNode > max {
		max = curLevelNode
	}
	return max
}

func main() {
	node6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		Val:   4,
		Left:  node6,
		Right: node5,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  node3,
		Right: nil,
	}

	root := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node4,
	}
	/*
		二叉树为
				 1
		     2        4
		  3	       6      5
	*/
	//preOrder(root)\
	fmt.Println(maxWidth(root))
}