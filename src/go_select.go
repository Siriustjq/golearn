package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		time.Sleep(1000000)
		ch1 <- false
	}()

	go func() {
		time.Sleep(2000000)
		ch2 <- true
	}()

	select {
	//select是阻塞的，如果所有case的io条件都未被触发，将会一直阻塞
	case <-ch1:
		fmt.Println("失败！")
	case <-ch2:
		fmt.Println("成功！")
	}
}
