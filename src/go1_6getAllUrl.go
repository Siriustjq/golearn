package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //启动一个goroutine去处理相关业务逻辑
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //从通道中读取数据，通道是goroutine之间通信的工具
	}
	close(ch) //防止资源泄漏
	fmt.Printf("%2.fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //将错误输出至ch通道中
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) //不需要http.Get(url)获得的信息，只是记录时间而已
	resp.Body.Close()                                 //避免泄露资源
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs,%7d,%s", secs, nbytes, url)
}
