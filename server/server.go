package main

import (
	"fmt"
	"net"
	"strconv"
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
	isEmpty bool
}

var start bool

type newGame interface {
	checkerboard()
}

var table [10][9]Piece

func checkerboarder() {
	start = true
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			table[i][j].isEmpty = true
		}
	}

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
	str = "〇一二三四五六七八\n"
	for i := 0; i < 10; i++ {

		for j := 0; j < 9; j++ {
			if table[i][j].isAlive == false {
				str += "十"
				// fmt.Print("十")
			} else if table[i][j].isEmpty {
				str += "十"
			} else {
				str += table[i][j].name
				// fmt.Print(table[i][j].name)
			}
		}
		// fmt.Println()
		s := strconv.Itoa(i)
		str += s
		str += "\n"
	}
	str += "〇一二三四五六七八\n"
	return str
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(x1, y1, x2, y2 int) bool {
	ret := true
	t := table[x1][y1].type1
	if x1 == x2 && y1 == y2 {
		return false
	}
	if x2 > 9 || x2 < 0 || y2 > 8 || y2 < 0 {
		return false
	}
	if table[x1][y1].isEmpty {
		return false
	}
	if table[x2][y2].isEmpty == false && table[x1][y1].camp == table[x2][y2].camp {
		return false
	}
	if t == rook {
		if x1 != x2 && y1 != y2 {
			return false
		}
		if x1 == x2 {
			if y1 > y2 {
				t := y1
				y1 = y2
				y2 = t
			}
			for i := y1 + 1; i < y2; i++ {
				if table[x1][i].isEmpty == false {
					return false
				}
			}
		} else if y1 == y2 {
			if x1 > x2 {
				t := x1
				x1 = x2
				x2 = t
			}
			for i := x1 + 1; i < x2; i++ {
				if table[i][y1].isEmpty == false {
					return false
				}
			}
		}
	} else if t == knight {
		if x1 == x2 {
			return false
		}
		if y1 == y2 {
			return false
		}
		if (abs(x1-x2) + abs(y1-y2)) != 3 {
			return false
		}
		if abs(y1-y2) == 2 {
			return table[x1][(y1+y2)/2].isEmpty
		}
		if abs(x1-x2) == 2 {
			return table[(x1+x2)/2][y1].isEmpty
		}
	} else if t == bishop {
		if abs(x1-x2) != 2 || abs(y1-y2) != 2 {
			return false
		}
		return table[(x1+x2)/2][(y1+y2)/2].isEmpty
	} else if t == escort {
		if abs(x1-x2) != 1 || abs(y1-y2) != 1 {
			return false
		}
		if table[x1][y1].camp == 0 {
			if x2 > 2 || y2 < 3 || y2 > 5 {
				return false
			}
		} else if table[x1][y1].camp == 1 {
			if x2 < 7 || y2 < 3 || y2 > 5 {
				return false
			}
		}
	} else if t == pawn {
		if (abs(x1-x2) + abs(y1-y2)) > 1 {
			return false
		}
		c := table[x1][y1].camp
		if c == 0 {
			if x2 < x1 {
				return false
			}
			if x2 <= 4 && y1 != y2 {
				return false
			}
		} else if c == 1 {
			if x2 > x1 {
				return false
			}
			if x2 >= 5 && y1 != y2 {
				return false
			}
		}
	} else if t == cannon {
		if x2 > 9 {
			return false
		}
		if x1 > 9 {
			return false
		}
		if x1 != x2 && y1 != y2 {
			return false
		}
		if x1 == x2 {
			sum := 0
			if y1 > y2 {
				p := y1
				y1 = y2
				y2 = p
			}

			for i := y1 + 1; i < y2; i++ {
				if table[x1][i].isEmpty == false {
					sum++
				}
			}
			if sum > 1 {
				return false
			} else if sum == 1 {
				if table[x1][y1].isEmpty == true {
					return false
				}
				if table[x1][y2].isEmpty == true {
					return false
				}
			} else {
				if table[x1][y1].isEmpty == false {
					return false
				}
				if table[x1][y2].isEmpty == false {
					return false
				}
			}
		} else {
			sum := 0
			if x1 > x2 {
				p := x1
				x1 = x2
				x2 = p
			}

			for i := x1 + 1; i < x2; i++ {
				if table[i][y1].isEmpty == false {
					sum++
				}
			}
			if sum > 1 {
				return false
			} else if sum == 1 {
				if table[x1][y1].isEmpty == true {
					return false
				}
				if table[x1][y2].isEmpty == true {
					return false
				}
			} else {
				if table[x1][y1].isEmpty == false {
					return false
				}
				if table[x1][y2].isEmpty == false {
					return false
				}
			}
		}

	} else if t == general {
		if y2 > 5 {
			return false
		}
		if y2 < 3 {
			return false
		}
		if y1 > 5 {
			return false
		}
		if y1 < 3 {
			return false
		}

		if x1 != x2 && y1 != y2 {
			return false
		}
		if table[x1][y1].camp == 0 {
			if x2 > 2 {
				return false
			}
			if x1 > 2 {
				return false
			}
		} else {
			if x2 < 7 {
				return false
			}
			if x1 < 7 {
				return false
			}
		}
		if y1 > y2 {
			p := y1
			y1 = y2
			y2 = p
		}
		if x1 > x2 {
			p := x1
			x1 = x2
			x2 = p
		}
		if y2-y1 != 1 {
			return false
		}
		if x2-x1 != 1 {
			return false
		}

	}
	return ret
}

func winner() int {
	var g [2]bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j].isEmpty {
				continue
			}
			if table[i][j].type1 == general {
				g[table[i][j].camp] = true
			}
		}
	}
	if g[0] && g[1] {
		return -1
	} else if g[0] == false && g[1] {
		return 1
	} else if g[1] == false && g[0] {
		return 0
	} else {
		fmt.Printf("error!!!!!!\n")
		return -2
	}
}

func move(from, to string) {

	x1, err := strconv.Atoi(from[0:1])
	y1, err := strconv.Atoi(from[1:2])
	x2, err := strconv.Atoi(to[0:1])
	y2, err := strconv.Atoi(to[1:2])
	if err != nil {
		fmt.Println("error")
	}
	table[x1][y1].isEmpty = true
	table[x2][y2].camp = table[x1][y1].camp
	table[x2][y2].isAlive = table[x1][y1].isAlive
	table[x2][y2].isEmpty = false
	table[x2][y2].name = table[x1][y1].name
	table[x2][y2].type1 = table[x1][y1].type1
	table[x2][y2].x = x2
	table[x2][y2].y = y2
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
		data := make([]byte, 1024)
		total, err := tcpConn.Read(data)
		str := string(data[:total])
		nickname := strings.Split(str, ":")[0]
		nickname = nickname[1:(len(nickname) - 2)]
		command := strings.Split(str, " ")[1][0:5]
		eligible := true
		fmt.Printf("===" + command + "===\n")
		if strings.Compare(command, "chess") == 0 { //开始新游戏
			checkerboarder()
			fmt.Println(tostring(), err) //打印到server屏幕
		} else if strings.Compare(command, "/move") == 0 {
			from := strings.Split(str, " ")[2][0:2]
			to := strings.Split(str, " ")[3][0:2]
			//check
			x1, err := strconv.Atoi(from[0:1])
			y1, err := strconv.Atoi(from[1:2])
			x2, err := strconv.Atoi(to[0:1])
			y2, err := strconv.Atoi(to[1:2])
			if err != nil {
				fmt.Println("error")
			}
			eligible = check(x1, y1, x2, y2)
			if eligible && start {
				move(from, to)
			}

			fmt.Println(tostring(), err)
		} else if strings.Compare(command, "/load") == 0 {

		} else if strings.Compare(command, "/save") == 0 {

		}
		fmt.Println(string(data[:total]), err) //打印到server屏幕
		flag := checkErr(err)
		if flag == 0 {
			break
		}

		for _, conn := range ConnMap {
			if conn.RemoteAddr().String() == tcpConn.RemoteAddr().String() && strings.Compare(command, "chess") != 0 && strings.Compare(command, "/move") != 0 {
				continue
			}
			if strings.Compare(command, "chess") == 0 {
				conn.Write([]byte(tostring()))
			} else if strings.Compare(command, "/move") == 0 {
				if start == false {
					conn.Write([]byte("The game is finished!\n"))
					continue
				}
				if eligible == false {
					conn.Write([]byte("Eligible move!\n"))
					continue
				}
				if winner() == 0 {
					conn.Write([]byte("Red is the winner\n" + tostring()))
				} else if winner() == 1 {
					conn.Write([]byte("Black is the winner\n" + tostring()))
				} else {
					conn.Write([]byte(tostring()))
				}
			} else {
				conn.Write(data[:total])
			}

		}
		if winner() != -1 {
			start = false
		}
	}
}

func main() {
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
