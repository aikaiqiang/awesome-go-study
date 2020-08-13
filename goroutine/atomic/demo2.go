package main

import (
	"fmt"
	"runtime"
	"time"
)

func printTime(n int) {
	now := time.Now()
	fmt.Printf("Index: %d, Second: %d, NanoSecond: %d \n", n, now.Second(), now.Nanosecond())
}

func main() {
	runtime.GOMAXPROCS(1) // 设置使用一个物理核
	go func() {
		printTime(2)
		panic("hello goroutine")
	}()
	printTime(1)
	sum := 0
	for i := 0; i < 666666666; i++ {
		sum += i
	}
	fmt.Printf("sum : %d \n", sum)
}
