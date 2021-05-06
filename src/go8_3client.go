package main

//仿照netcat构造一个简单的客户端

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Print("Connection has been refused ! ")
	}
	defer conn.Close()
	//并发接受聊天信息
	go messcopy(os.Stdout, conn)
	messcopy(conn, os.Stdin)
}

func messcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
