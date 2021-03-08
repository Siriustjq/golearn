package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	time.Sleep(time.Second)
}

func test() {
	ch := make(chan int, 4)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		close(ch)
	}()

	go func() {
		for a := range ch {
			fmt.Print(a)
		}
	}()

	fmt.Print("test is over")
}
