package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	odd := make(chan struct{})
	even := make(chan struct{})
	go printEven(wg, odd, even)
	go printOdd(wg, odd, even)
	wg.Wait()
}

func printOdd(wg *sync.WaitGroup, odd, even chan struct{}) {
	defer wg.Done()
	for i := 1; i < 100; i = i + 2 {
		fmt.Println(i)
		// 通知偶数可以执行
		odd <- struct{}{}
		// 等待偶数打印完毕
		<-even
	}
}

func printEven(wg *sync.WaitGroup, odd, even chan struct{}) {
	defer wg.Done()
	for i := 2; i <= 100; i = i + 2 {
		// 等待奇数打印完成
		<-odd
		// 开始打偶数
		fmt.Println(i)
		// 通知奇数执行
		even <- struct{}{}
	}
}
