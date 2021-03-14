package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	f1 := squares()
	fmt.Println(f1())
}

//这里相当于定义了一个函数变量，GO语言中称之为函数字面量
func squares() func() int {
	var x int
	fmt.Printf("此时x的值为：%d\n", x)
	var fre int
	//以上为类似的初始化过程，在函数字面量被第一次赋值时运行
	return func() int {
		fre++
		fmt.Printf("这第%d次调用\n", fre)
		x++
		return x * x
	}
}
