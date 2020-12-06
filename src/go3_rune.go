package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	testStr := "this is a test!"
	testStr1 := "这是一个测试"
	fmt.Println("The length of str is: ", len(testStr))
	//fmt.Println("字符串长度为: ",len(testStr1))
	fmt.Println("字符串长度为: ", utf8.RuneCountInString(testStr1))
	fmt.Println("字符串长度为: ", len([]rune(testStr1)))
}

/**
这里主要是来测试验证rune数据类型的作用
主要是来表示unicode或者utf-8字符
*/
//fmt.Println("字符串长度为: ",utf8.RuneCountInString(testStr1))
