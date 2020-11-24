package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //http发出请求之后会返回一个结果封装在resp中，另外还将返回错误信息
		if err != nil {
			fmt.Fprintf(os.Stderr, "getUrl %v\n", err)
			os.Exit(1) //错误反生后，进程退出时返回的状态码
		}
		_, err = io.Copy(os.Stdout, resp.Body) //将从resp.Body中读数据内容读到Stdout中
		//直接将内容读到os.Stdout,这里可以有多种输出方式
		//例如：./go1_5test1.7 http://www.baidu.com >baidu.txt
		//上述指令的就是将结果直接写入到了baidu.txt文件中
		//./go1_4gif >out.gif 则是直接输出一个gif文件
		if err != nil {
			fmt.Fprintf(os.Stderr, "getUrl is fail, the content is :%v\n", err)
			os.Exit(1)
		}
	}
}
