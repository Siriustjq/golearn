package main

/**
主要涉及到go语言中的迭代变量的捕获
*/

import (
	"fmt"
)

func main() {
	var handles []func()
	var numsSilces []int
	for i := 0; i < 10; i++ {
		numsSilces = append(numsSilces, i)
	}
	for _, num := range numsSilces {
		num := num
		handles = append(handles, func() {
			fmt.Println(num)
		})
	}

	for _, handle := range handles {
		handle()
	}
}
