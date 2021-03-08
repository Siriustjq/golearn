package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func main() {
	time.Sleep(1 * time.Second)
	go dosome()
	time.Sleep(10 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	fmt.Println("cancelled")
	time.Sleep(2 * time.Second)
}

func dosome() {
	for {
		if done {
			return
		}
		fmt.Println("tick")
		time.Sleep(1 * time.Second)
	}
}
