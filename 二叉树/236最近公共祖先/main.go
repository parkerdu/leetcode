package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, father map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		father[root.Left] = root
	}
	if root.Right != nil {
		father[root.Right] = root
	}
	preOrder(root.Left, father)
	preOrder(root.Right, father)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	father := make(map[*TreeNode]*TreeNode)
	father[root] = root
	path := make(map[*TreeNode]struct{})
	// step1: 前序遍历树，存储父亲节点
	preOrder(root, father)
	// step2: p向上走，走过的节点记录下来
	for p != father[p] {
		path[p] = struct{}{}
		p = father[p]
	}
	// step3: q向上走，走的路在p的路径上则为lca
	for q != father[q] {
		_, ok := path[q]
		if ok {
			return q
		}
		q = father[q]
	}
	return p
}

// 法二：递归
/*
定义该函数是递归的关键
若p,q都在root下面，返回公共祖先
若p,q只有一个在root下面，返回p或者q
若p,q都不在root下，返回nil
*/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 当前节点为p或者q返回
	if root == p || root == q {
		return root
	}
	// 左子树找p,q的公共祖先，p,q都不在左边返回nil, p,q有一个在里面返回p或q, p,q都在里面返回公共祖先
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	// 左右都不空，说明p,q一个在左，一个在右，返回root
	if left != nil && right != nil {
		return root
	}
	// 左不空，右空，说明p,q都在左，此时left就是最近公共祖先
	if left != nil {
		return left
	}
	return right
}

func main() {
	// 构造，1,2,3，nil,nil,4,5
	/*
					 1
				   2   3
			          4  5
		             6
	*/
	six := &TreeNode{Val: 6}
	five := &TreeNode{
		Val: 5,
	}
	four := &TreeNode{Val: 4, Left: six}
	three := &TreeNode{
		Val:   3,
		Left:  four,
		Right: five,
	}
	two := &TreeNode{
		Val: 2,
	}
	root := &TreeNode{Val: 1, Left: two, Right: three}
	lowestCommonAncestor1(root, two, six)
}
