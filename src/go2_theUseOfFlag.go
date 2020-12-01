package main

import (
	"flag"
	"fmt"
)

/**
通过命令行参数输入姓名、年龄、性别，并输出
*/
//自定义命令行参数
var name = flag.String("name", "tjq", "the name of person")
var age = flag.Int("age", 18, "the age of person")
var gender = flag.String("gender", "male", "the gender of person")

func main() {
	flag.Parse() //解析
	fmt.Printf("name = %s, age = %d, gender = %s", *name, *age, *gender)
	fmt.Println()
	fmt.Println("未被解析的命令行为：", flag.Args())
}
