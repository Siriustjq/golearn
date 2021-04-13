package main

import "fmt"

func randomRes() bool {
	return false
}

func main() {
	//switch是不会阻塞的，如果没有任何一个分支满足条件，直接跳出选择体了
	switch randomRes() {
	case true:
		fmt.Println("the result is true")
	case false:
		fmt.Println("the result is false")
	}

	switch randomRes(); { //注意这里的换行问题，这里换行了的话，相当于加了一个分号，那么相当于这个switch判断已经没有条件了，自然将会输出true
	case true:
		fmt.Println("the result is true")
	case false:
		fmt.Println("the result is false")
	}
}
