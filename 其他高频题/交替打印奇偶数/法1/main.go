package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁 + chan 行不通
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	mutex := &sync.Mutex{}
	even := make(chan struct{})
	go printEven(wg, mutex, even)
	go printOdd(wg, mutex, even)
	wg.Wait()
}

func printOdd(wg *sync.WaitGroup, m *sync.Mutex, even chan struct{}) {
	defer wg.Done()
	for i := 1; i < 100; i = i + 2 {
		// 获取锁，打印
		m.Lock()
		fmt.Println(i)
		m.Unlock()
		// 等待偶数拿到锁且打印完毕
		<-even
	}
}

func printEven(wg *sync.WaitGroup, m *sync.Mutex, even chan struct{}) {
	defer wg.Done()
	for i := 2; i <= 100; i = i + 2 {
		// 等待奇数打印完成

		// 开始打偶数
		m.Lock()
		fmt.Println(i)
		m.Unlock()
		even <- struct{}{}
		time.Sleep(time.Second)
	}
}
