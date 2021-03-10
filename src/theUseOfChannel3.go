package main

import "fmt"

func main() {
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(x int) {
			sendRpc2(x)
			done <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-done
	}
}

func sendRpc2(x int) {
	fmt.Println(x)
}
