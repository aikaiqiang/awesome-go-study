package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type Person struct {
	Money int64
}

func main() {
	runtime.GOMAXPROCS(1) //配置只使用一个物理核
	p := Person{Money: 100}
	go func() {
		atomic.AddInt64(&p.Money, 1000)
	}()
	atomic.AddInt64(&p.Money, 1000)
	fmt.Printf("Money: %d\n", atomic.LoadInt64(&p.Money))
}
