package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	count := 0
	finshed := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu) //设置条件变量，提高程序性能
	rand.Seed(time.Now().UnixNano())

	//假设一共10个节点
	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finshed++
			cond.Broadcast()
		}()
	}
	//for {
	//	mu.Lock()
	//	if count >5 || finshed == 10 {
	//		break
	//	}
	//	mu.Unlock()
	//}
	mu.Lock()
	for count < 5 && finshed != 10 {
		cond.Wait()
	}
	if count > 5 {
		fmt.Println("received major voted!")
	} else {
		fmt.Println("failed!")
	}
	mu.Unlock()
}

func requestVote() bool {
	//随机模拟投票时间和投票结果
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
