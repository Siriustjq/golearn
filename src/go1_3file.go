package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	//从命令行中获取参数，这里指的是文件列表，如果文件为空，那么就从标准输入中获取内容
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			//文件打开会返回两个参数，一个文件本身，另外一个是错误信息
			f, err := os.Open(arg)
			//错误信息如果等于nil，那么说明读取正常，反之发生错误，需要使用Fprintf输出错误
			if err != nil {
				fmt.Fprintf(os.Stderr, "go1_3file.go%t", err)
			}
			countLines(f, counts)
			//文件使用完毕一定要注意close
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
