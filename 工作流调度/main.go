package main

import (
	"fmt"
	"sync"
)

/*
工作流调度，支持正常dag调度，支持添加条件边，支持有环图调度
*/

const END = "_END_"

type Handler func(id string) (out interface{})

// 根据当前节点的输出，决定走哪一条或者多条边，相当于一个过滤器
type Router func(out interface{}) []string

type Node struct {
	Id       string
	In       int // 入度, 当前节点成功，下面所有孩子节点的入度-1
	Out      int // 出度
	Parent   []*Node
	Children map[string]*Node // todo 用集合是个更好的选择
	Handle   Handler
	Route    Router
	Err      error
	Success  bool
	output   interface{} // 该节点输出结果
}

// 封装一个函数，用来分辨条件边还是正常边
func (n *Node) acquireChild(out interface{}) map[string]*Node {
	if n.Route != nil { // 该节点下面是条件边，使用用户定义的route
		children := make(map[string]*Node)
		for _, id := range n.Route(out) {
			children[id] = n.Children[id]
		}
		return children
	}
	return n.Children // 该节点下面是正常边，直接返回child
}

// 从当前node开始向上找祖宗节点，如果能遇到root说明是形成了环
func (n *Node) isCycle(root *Node) bool {
	for i := range n.Parent {
		if n.Parent[i] == root || n.Parent[i].isCycle(root) {
			return true
		}
	}
	return false
}

func (n *Node) Start(wg *sync.WaitGroup) {
	// step1:检查父节点完成情况,这种通过入度来检查是否ready可以少一次遍历
	// 为了适应环节点，当入度不为0时，还要继续查看当前未完成的父亲节点是不环造成的
	// 对于一个条件边的父亲来说，孩子是不定的，但是从孩子的角度来看，父亲是唯一的，所以若是当前的节点的父节点一直往上找能找到当前节点，说明是环节点可以略过
	//if n.In != 0 {
	//	return
	//}
	defer wg.Done()
	for i := range n.Parent {
		if !n.Parent[i].Success && !n.isCycle(n) {
			return
		}
	}
	// step2: 执行自身逻辑
	if n.Id == END {
		fmt.Println("执行end")
		return
	}
	n.output = n.Handle(n.Id)
	n.Success = true
	// step3: 根据当前输出，过滤孩子节点，更新孩子节点入度
	// step3: 调度孩子节点
	for _, child := range n.acquireChild(n.output) {
		child.In--
		wg.Add(1)
		go child.Start(wg)
	}
}

type edge struct {
	From *Node
	To   *Node
}

type Workflow struct {
	root  []*Node
	Nodes map[string]*Node
}

func (w *Workflow) Schedule() {
	wg := &sync.WaitGroup{}
	for i := range w.root {
		wg.Add(1)
		go w.root[i].Start(wg)
	}
	wg.Wait()
}

func (w *Workflow) AddNode(id string, handle Handler) error {
	_, ok := w.Nodes[id]
	if ok {
		return fmt.Errorf("id exist")
	}
	w.Nodes[id] = &Node{
		Id:       id,
		Handle:   handle,
		Parent:   make([]*Node, 0),
		Children: map[string]*Node{},
	}
	return nil
}

func (w *Workflow) AddEdge(from, to string) error {
	fNode, ok := w.Nodes[from]
	if !ok {
		return fmt.Errorf("%s not exist", from)
	}
	tNode, ok := w.Nodes[to]
	if !ok {
		return fmt.Errorf("%s not exist", to)
	}
	fNode.Children[to] = tNode
	tNode.Parent = append(tNode.Parent, fNode)
	return nil
}

func (w *Workflow) AddConditionEdge(from string, route Router, to []string) error {
	// todo golang能够直接获取route的返回值列表，类似python 的Literal["multiply", "__end__"]:
	fNode, ok := w.Nodes[from]
	if !ok {
		return fmt.Errorf("%s not exist", from)
	}
	fNode.Route = route
	for i := range to {
		tNode, ok := w.Nodes[to[i]]
		if !ok {
			return fmt.Errorf("%s not exist", to)
		}
		fNode.Children[to[i]] = tNode
		tNode.Parent = append(tNode.Parent, fNode)
	}
	return nil
}

func (w *Workflow) EntryPoint(root []string) error {
	for i := range root {
		node, ok := w.Nodes[root[i]]
		if !ok {
			return fmt.Errorf("%s not exist", root[i])
		}
		w.root = append(w.root, node)
	}
	return nil
}

func NewWorkflow() *Workflow {
	return &Workflow{
		root:  make([]*Node, 0),
		Nodes: make(map[string]*Node),
	}
}

func main() {
	w := NewWorkflow()
	h := func(id string) interface{} {
		fmt.Println(id)
		return nil
	}
	w.AddNode("a", h)
	w.AddNode("b", h)

	handleC := func(id string) interface{} {
		fmt.Println(id)

		return 100
	}
	w.AddNode("c", handleC)

	w.AddEdge("a", "b")
	w.AddEdge("b", "c")
	r := func(out interface{}) []string {
		// c节点输出>=100 则结束，否则走b
		val, _ := out.(int)
		if val >= 100 {
			return []string{END}
		}
		return []string{"b"}
	}
	w.AddConditionEdge("c", r, []string{"b", END})
	w.EntryPoint([]string{"a"})
	w.Schedule()
}
