package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
创建一个服务，返回请求的url地址，即url.path
运行该服务器程序只需要在命令行输入：
go run go1_7server1.go &
千万不要忘记后面的"&"
*/

func main() {
	http.HandleFunc("/", handler) //指定请求拦截后调用的处理函数
	log.Fatal(http.ListenAndServe("localhost:8003", nil))
}

//用于处理返回请求url路径信息的函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
