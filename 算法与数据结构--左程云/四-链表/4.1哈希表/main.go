package main

import "fmt"

type student struct {
	age int
	class int
}

func main() {
	s1 := student{
		age: 1,
		class: 1,
	}
	m := make(map[student]int)
	m[s1] = 1
	for k, _ := range m{
		fmt.Printf("key 的内容用地址表示：%p", k)

	}
	fmt.Println(m[s1])
}
