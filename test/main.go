package main

import (
	"fmt"
	"sync"
)

func main()  {
	// channel 带缓冲通道
	cNumbs := make(chan int, 3)
	fmt.Println("通道长度:", len(cNumbs))
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		for{
			fmt.Println("开始读取")
			select {
			case b := <-cNumbs:
				fmt.Println("通道内容读取:", b)
			}
		}
	}()
	wg.Wait()
}