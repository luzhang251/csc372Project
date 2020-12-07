/*
//https://docs.google.com/document/d/1aadvk1rOnTXXjB7WmXU_0qmTFzRn_LfsjvJCsCaGMhQ/edit?usp=sharing
Assignment: Final Project - Part 3 Creative Program
Author: Lu Zhang, Zhenyu Yuan
Course: CSc 372
Instructor: L. McCann
TA(s): Tito Ferra and Josh Xiong
Due Date: Dec 7, 2020
Description: A Chinese Chess Game
			 This is the code of client-side
			 basicly this client-side only
			 contains TCPconn structure code
			 and features of send and receive messages.
			 Golang manual contains a good tutorial of the use of
			 net package.

Language: Golang
Ex. Packages: None.(all original packages, no 3rd party packages)
Deficiencies: None.
*/
package main

// import packages
import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// check error
var ch chan int = make(chan int)
var nickname string

// read the content that is sended from the server
func reader(conn *net.TCPConn) {
	buff := make([]byte, 50240)
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
	// conn is the connection buffer
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Server is not starting")
		os.Exit(0)
	}
	defer conn.Close()
	go reader(conn)
	// get the nickname of the user
	fmt.Println("Enter your nickname:")
	fmt.Scanln(&nickname)
	fmt.Println("Your nickname is", nickname)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		// send message (command) to the server
		b := []byte("<" + nickname + ">" + ": " + msg + "\n")
		conn.Write(b)
		select {
		case <-ch:
			fmt.Println("Server ERROR")
			os.Exit(1)
		default:
		}
	}
}
