package main

import (
	"fmt"
	"io/ioutil"
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
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "getUrl is fail, the content is : %s:%v\n", b, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
