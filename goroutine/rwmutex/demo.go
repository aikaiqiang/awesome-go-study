package main

import (
	"fmt"
	"sync"
)

// Person struct
type Person struct {
	Money   int
	rwMutex sync.RWMutex
}

// GetMoney 获取用户金钱
func (p *Person) GetMoney() int {
	p.rwMutex.RLock()
	defer p.rwMutex.RUnlock()
	money := p.Money
	return money
}

// AddMoney 设置用户金钱
func (p *Person) AddMoney(diff int) {
	p.rwMutex.Lock()
	defer p.rwMutex.Unlock()
	p.Money += diff
}

// 主函数
func test() {
	p := Person{Money: 100}
	go func() {
		p.AddMoney(1000)
	}()
	fmt.Printf("Money: %d\n", p.GetMoney())
}
