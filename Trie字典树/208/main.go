package main

import "fmt"

/*
2024/04/27 夜里2点22分，时隔5年再次遇到该题，还是不会
19年听覃超，24年听左程云
 */

type Trie struct {
	edges [26]*Trie
	end int
}

// 值调用指针方法会发生什么？答：编译器会自动加上地址来调用
func Constructor() Trie {
	return Trie{}
}


func (t *Trie) Insert(word string)  {
	// w是rune类型？
	cur := t
	for _,w := range word {
		index := w - 'a'
		if cur.edges[index] == nil {
			// 没有该字符新建
			cur.edges[index] = &Trie{}
		}
		// 更新cur为下一个
		cur = cur.edges[index]
	}
	cur.end++
}


func (t*Trie) Search(word string) bool {
	// 遍历字母，看cur的index是否为空，空则表示不存在，循环走完表示存在
	cur := t
	for _,w := range word {
		index := w - 'a'
		if cur.edges[index] == nil {
			return false
		}
		cur = cur.edges[index]
	}
	if cur.end != 0 {
		return true
	}
	return false
}


func (t *Trie) StartsWith(prefix string) bool {
	// 遍历字母，看cur的index是否为空，空则表示不存在，循环走完表示存在
	cur := t
	for _,w := range prefix {
		index := w - 'a'
		if cur.edges[index] == nil {
			return false
		}
		cur = cur.edges[index]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()
	obj.Insert("abc")
	obj.Insert("abd")
	fmt.Println(obj.Search("abc"))
	fmt.Println(obj.StartsWith("bc"))


}