package main

//实现一个并发的回声服务器（该并发是从两个角度展开的，多个连接并发进行，同一个连接内多次并发请求）

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
		}
		go handleConn2(conn)
	}
}

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
