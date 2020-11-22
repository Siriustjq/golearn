package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
本程序，主要是从控制台输入，并通过map输出每次输入的重复行
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	//注意，此处忽略了input.err可能发生的错误
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
