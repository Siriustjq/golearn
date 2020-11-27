package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/**
创建一个返回URL.Path的服务器，并且同时可以记录访问次数，即同时还是一个计数器服务器
*/
var lockTool sync.Mutex //这是一个同步工具
var count int

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8007", nil))
}

//回升服务器处理逻辑，添加同步选项
func handler1(w http.ResponseWriter, r *http.Request) {
	lockTool.Lock() //使其在同一时间内只有一个goroutine在运行
	count++
	lockTool.Unlock()
	fmt.Fprintf(w, "URL.Path is %q\n", r.URL.Path)
}

//回显目前为止访问次数
func counter(w http.ResponseWriter, r *http.Request) {
	lockTool.Lock() //使其在同一时间内只有一个goroutine在运行
	fmt.Fprintf(w, "Count %d\n", count)
	lockTool.Unlock()
}
