package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Connected has been refused！", err)
	}
	defer conn.Close()
	go mesCopy(os.Stdout, conn)
	mesCopy(conn, os.Stdin)
}

func mesCopy(des io.Writer, res io.Reader) {
	if _, err := io.Copy(des, res); err != nil {
		log.Print("wrong!")
	}
}
