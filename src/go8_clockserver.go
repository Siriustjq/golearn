package main

/**
利用go天然的并发机制实现一个时间服务器，实现多客户端的并发访问，每隔1s返回客户端一个时间
*/

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Print("Something wrong！Connecting failed ！")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Something wrong！Connecting failed ！")
			continue
		}
		//调用具体每隔1s返回时间的程序逻辑
		//每一个连接都用一个goroutine去处理
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print("handleConn wrong!")
			return
		}
		//休息1s钟
		time.Sleep(1 * time.Second)
	}
}
