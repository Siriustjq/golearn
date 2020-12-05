package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, num := range os.Args[1:] {
		//strconv 包可以将字符串类型转换为其他类型
		num, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("something wrong")
		} else {
			fmt.Println(num)
			fmt.Println("the number of one is ", numberOfOne(num))
		}
	}
}

func numberOfOne(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
