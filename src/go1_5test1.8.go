package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") { //增强版，首先判断是否具有http前缀，没有就先加上再去处理url
			url = "http://" + url
		}
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
