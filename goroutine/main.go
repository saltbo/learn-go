package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func test() (interface{}, error) {
	var (
		routines = 3
		errChan  = make(chan error, routines)
	)

	// 总订单数
	go func() {
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		errChan <- nil
	}()

	// 总时长
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		errChan <- fmt.Errorf("error happened!")
	}()

	// 关联商品数
	go func() {
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		errChan <- nil
	}()

	fmt.Println(runtime.NumGoroutine())
	for i := 0; i < routines; i++ {
		if err := <-errChan; err != nil {
			return nil, err
		}
	}
	close(errChan)
	return 111, nil
}

func main() {
	_, err := test()
	fmt.Println(err)
	fmt.Println(runtime.NumGoroutine())

	for range time.Tick(time.Second) {
		fmt.Println(runtime.NumGoroutine())
	}
}
