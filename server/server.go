package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	// pieces
	general = 0
	rook    = 1
	knight  = 2
	cannon  = 3
	bishop  = 4
	escort  = 5
	pawn    = 6

	general0 = "帅"
	rook0    = "车"
	knight0  = "马"
	cannon0  = "炮"
	bishop0  = "相"
	escort0  = "仕"
	pawn0    = "兵"

	general1 = "将"
	rook1    = "車"
	knight1  = "馬"
	cannon1  = "砲"
	bishop1  = "象"
	escort1  = "士"
	pawn1    = "卒"
	// for camp
	red   = 0
	black = 1
)

type Piece struct {
	camp    int
	name    string
	type1   int
	isAlive bool
	x       int
	y       int
}

type newGame interface {
	checkerboard()
}

var table [10][9]Piece

func checkerboarder() {
	table[0][0] = Piece{camp: 0, type1: rook, isAlive: true, name: rook0}
	table[0][1] = Piece{camp: 0, type1: knight, isAlive: true, name: knight0}
	table[0][2] = Piece{camp: 0, type1: bishop, isAlive: true, name: bishop0}
	table[0][3] = Piece{camp: 0, type1: escort, isAlive: true, name: escort0}
	table[0][4] = Piece{camp: 0, type1: general, isAlive: true, name: general0}
	table[0][5] = Piece{camp: 0, type1: escort, isAlive: true, name: escort0}
	table[0][6] = Piece{camp: 0, type1: bishop, isAlive: true, name: bishop0}
	table[0][7] = Piece{camp: 0, type1: knight, isAlive: true, name: knight0}
	table[0][8] = Piece{camp: 0, type1: rook, isAlive: true, name: rook0}
	table[2][1] = Piece{camp: 0, type1: cannon, isAlive: true, name: cannon0}
	table[2][7] = Piece{camp: 0, type1: cannon, isAlive: true, name: cannon0}
	table[3][0] = Piece{camp: 0, type1: pawn, isAlive: true, name: pawn0}
	table[3][2] = Piece{camp: 0, type1: pawn, isAlive: true, name: pawn0}
	table[3][4] = Piece{camp: 0, type1: pawn, isAlive: true, name: pawn0}
	table[3][6] = Piece{camp: 0, type1: pawn, isAlive: true, name: pawn0}
	table[3][8] = Piece{camp: 0, type1: pawn, isAlive: true, name: pawn1}
	table[6][0] = Piece{camp: 1, type1: pawn, isAlive: true, name: pawn1}
	table[6][2] = Piece{camp: 1, type1: pawn, isAlive: true, name: pawn1}
	table[6][4] = Piece{camp: 1, type1: pawn, isAlive: true, name: pawn1}
	table[6][6] = Piece{camp: 1, type1: pawn, isAlive: true, name: pawn1}
	table[6][8] = Piece{camp: 1, type1: pawn, isAlive: true, name: pawn1}
	table[7][1] = Piece{camp: 1, type1: cannon, isAlive: true, name: cannon1}
	table[7][7] = Piece{camp: 1, type1: cannon, isAlive: true, name: cannon1}
	table[9][0] = Piece{camp: 1, type1: rook, isAlive: true, name: rook1}
	table[9][1] = Piece{camp: 1, type1: knight, isAlive: true, name: knight1}
	table[9][2] = Piece{camp: 1, type1: bishop, isAlive: true, name: bishop1}
	table[9][3] = Piece{camp: 1, type1: escort, isAlive: true, name: escort1}
	table[9][4] = Piece{camp: 1, type1: general, isAlive: true, name: general1}
	table[9][5] = Piece{camp: 1, type1: escort, isAlive: true, name: escort1}
	table[9][6] = Piece{camp: 1, type1: bishop, isAlive: true, name: bishop1}
	table[9][7] = Piece{camp: 1, type1: knight, isAlive: true, name: knight1}
	table[9][8] = Piece{camp: 1, type1: rook, isAlive: true, name: rook1}
}

func tostring() string {
	var str string
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j].isAlive == false {
				str += "十"
				fmt.Print("十")
			} else {
				str += table[i][j].name
				fmt.Print(table[i][j].name)
			}
		}
		fmt.Println()
		str += "\n"
	}
	return str
}

var ConnMap map[string]*net.TCPConn

func checkErr(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("用户退出了")
			return 0
		}
		fmt.Println("错误")
		return -1
	}
	return 1
}

func say(tcpConn *net.TCPConn) {
	for {
		data := make([]byte, 128)
		total, err := tcpConn.Read(data)
		str := string(data[:total])
		fmt.Println("---" + str + "----\n")
		if strings.Compare(strings.Split(str, " ")[1], "chess") == 0 { //开始新游戏
			fmt.Println(strings.Split(str, " ")[1] + "++++++")
			tostring()
		}
		fmt.Println(string(data[:total]), err) //打印到server屏幕
		flag := checkErr(err)
		if flag == 0 {
			break
		}

		for _, conn := range ConnMap {
			if conn.RemoteAddr().String() == tcpConn.RemoteAddr().String() {
				continue
			}
			conn.Write(data[:total]) //发给用户
		}
	}
}

func main() {
	checkerboarder()
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	ConnMap = make(map[string]*net.TCPConn)
	for {
		tcpConn, _ := tcpListener.AcceptTCP()
		defer tcpConn.Close()

		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		fmt.Println("连接的客服端信息:", tcpConn.RemoteAddr().String())
		go say(tcpConn)
	}
}
