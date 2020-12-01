package main

import (
	"flag"
	"fmt"
	"strings"
)

/**
flag包主要是用来操作命令行参数的，可以自定义命令行参数，譬如flag.Bool以及flag.String等，返回值均为指针类型
*/
//使用 -n 忽略掉输出时末尾默认携带的换行符
var n = flag.Bool("n", false, "omit trailing newline")

//使用 -s seq 用seq替换掉输入参数之间默认的空格分隔符
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()                               //解析上面自定义的命令行参数
	fmt.Print(strings.Join(flag.Args(), *sep)) //flag.Args()返回的是没有被解析到的命令行字符串，即除自定义之外的命令行内容
	if !*n {
		fmt.Println()
	}
}
