package _map

import "sync"

// Golang 运行时明确禁止 map 的并发读写

// func main() {
//m := make(map[int]int)
//go func() {
//	for {
//		_ = m[1]
//	}
//}()
//
//go func() {
//	for {
//		m[2] = 1
//	}
//
//}()
//
//select {
//
//}

// updateDemo()
// }

func updateDemo() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	go func() {
		for {
			counter.RLock()
			_ = counter.m["some_key"]
			counter.RUnlock()
		}
	}()
	go func() {
		for {
			counter.Lock()
			counter.m["some_key"]++
			counter.Unlock()
		}
	}()
	select {}
}
