package main

import (
	"fmt"
	"time"
)

/**
该部分主要是去体会一下goruntine是并发执行的
*/

func main() {
	go spinner(100 * time.Millisecond) //该部分会一直运行，知道main运行结束，即斐波那契数列计算完毕
	const n = 45
	res := fabi(n)
	fmt.Printf("\r斐波那契数列在n=%d时，结果为%d", n, res) ///r就是回到行首的意思，相当于输入回车
}

//辅助验证函数spinner
func spinner(delay time.Duration) {
	wait := []string{0: "-", 1: "/", 2: "|", 3: "\\"}
	for {
		for _, r := range wait {
			fmt.Printf("\r%s", r)
			time.Sleep(delay)
		}
	}
}

//菲波那切数列
func fabi(n int) int {
	if n < 2 {
		return n
	}
	return fabi(n-1) + fabi(n-2)
}
