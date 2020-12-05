package main

import "fmt"

/**
主要是测试一下init初始化函数的作用
*/
var nums int

func init() {
	//init函数总是会按照顺序执行，无需进行任何显式调用
	nums = 15
}

func main() {
	fmt.Println(nums)
}
