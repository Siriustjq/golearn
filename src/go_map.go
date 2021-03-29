package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var testMap map[string]string

func main() {
	var strs []string
	testMap = make(map[string]string)
	for i := 1; i < 10; i++ {
		str := strconv.Itoa(i)
		str = "num" + str + ":"
		testMap[str] = strconv.Itoa(i)
	}
	fmt.Println(reflect.TypeOf(tools(testMap)))
	if tools(testMap) == "" {
		fmt.Println("空字符串")
	}
	if nil == strs {
		fmt.Println("空字符串数组")
	}
}

func tools(m map[string]string) string {
	//return m["num1:"]

	return m["num11:"]
}
