package main

import (
	"fmt"
	"net"
	"os"
)

var ch chan int = make(chan int)

var nickname string

func reader(conn *net.TCPConn) {
	buff := make([]byte, 128)
	for {
		j, err := conn.Read(buff)
		if err != nil {
			ch <- 1
			break
		}
		fmt.Println(string(buff[0:j]))
	}
}

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Server is not starting")
		os.Exit(0)
	}
	defer conn.Close()
	go reader(conn)
	fmt.Println("Enter your nickname:")
	fmt.Scanln(&nickname)
	fmt.Println("Your nickname is", nickname)

	for {
		var msg string
		fmt.Scanln(&msg)
		b := []byte("<" + nickname + ">" + ": " + msg)
		conn.Write(b)
		select {
		case <-ch:
			fmt.Println("Server ERROR")
			os.Exit(1)
		default:
		}
	}
}
