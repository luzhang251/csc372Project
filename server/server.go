package main

import (
    "fmt"
    "net"
)



var table [10][9]piece

const(
    // pieces
	general = 0
 	rook = 1
 	knight = 2
 	cannon = 3
	bishop = 4
 	escort = 5
	pawn = 6

    // for camp
    red = 0
    black = 1
)
    
type Piece struct {
    camp int
    type1  int
	isAlive bool
    x int
    y int
}

func newPiece(t, camp int) Piece{

}

type newGame interface{
    checkerboard()
}



func checkerboard(){
    //设置棋盘
	table[0][0] = Piece{camp: 0, type1: rook, isAlive: true}
	table[0][1] = Piece{camp: 0, type1: knight, isAlive: true}
	table[0][2] = Piece{camp: 0, type1: bishop, isAlive: true}
	table[0][3] = Piece{camp: 0, type1: escort, isAlive: true}
	table[0][4] = Piece{camp: 0, type1: general, isAlive: true}
	table[0][5] = Piece{camp: 0, type1: escort, isAlive: true}
	table[0][6] = Piece{camp: 0, type1: bishop, isAlive: true}
	table[0][7] = Piece{camp: 0, type1: knight, isAlive: true}
	table[0][8] = Piece{camp: 0, type1: rook, isAlive: true}

	table[2][1] = Piece{camp: 0, type1: cannon, isAlive: true}
	table[2][7] = Piece{camp: 0, type1: cannon, isAlive: true}

	table[3][0] = Piece{camp: 0, type1: pawn, isAlive: true}
	table[3][2] = Piece{camp: 0, type1: pawn, isAlive: true}
	table[3][4] = Piece{camp: 0, type1: pawn, isAlive: true}
	table[3][6] = Piece{camp: 0, type1: pawn, isAlive: true}
	table[3][8] = Piece{camp: 0, type1: pawn, isAlive: true}

	table[6][0] = Piece{camp: 1, type1: pawn, isAlive: true}
	table[6][2] = Piece{camp: 1, type1: pawn, isAlive: true}
	table[6][4] = Piece{camp: 1, type1: pawn, isAlive: true}
	table[6][6] = Piece{camp: 1, type1: pawn, isAlive: true}
	table[6][8] = Piece{camp: 1, type1: pawn, isAlive: true}

	table[7][1] = Piece{camp: 1, type1: cannon, isAlive: true}
	table[7][7] = Piece{camp: 1, type1: cannon, isAlive: true}

	table[9][0] = Piece{camp: 1, type1: rook, isAlive: true}
	table[9][1] = Piece{camp: 1, type1: knight, isAlive: true}
	table[9][2] = Piece{camp: 1, type1: bishop, isAlive: true}
	table[9][3] = Piece{camp: 1, type1: escort, isAlive: true}
	table[9][4] = Piece{camp: 1, type1: general, isAlive: true}
	table[9][5] = Piece{camp: 1, type1: escort, isAlive: true}
	table[9][6] = Piece{camp: 1, type1: bishop, isAlive: true}
	table[9][7] = Piece{camp: 1, type1: knight, isAlive: true}
	table[9][8] = Piece{camp: 1, type1: rook, isAlive: true}
}

func stringfyBoard() bytes.Buffer{
	var buf bytes.Buffer
	for i := 0; i < 10; i++{
		for j := 0; j < 9; j++{
			if(!table[i][j]) {
				buf.WriteString("十")
			} else if(table[i][j].type1 == rook && table[i][j].camp == 0){
				buf.WriteString("车")
			} else if(table[i][j].type1 == rook && table[i][j].camp == 1){
				buf.WriteString("車")
			} else if(table[i][j].type1 == rook && table[i][j].camp == 0){
				buf.WriteString("车")
			}else if(table[i][j].type1 == rook && table[i][j].camp == 0){
				buf.WriteString("车")
			}else if(table[i][j].type1 == rook && table[i][j].camp == 0){
				buf.WriteString("车")
			}else if(table[i][j].type1 == rook && table[i][j].camp == 0){
				buf.WriteString("车")
			}
		}
	}
}

var ConnMap map[string]*net.TCPConn

func checkErr(err error) int {
    if err != nil {
        if err.Error() == "EOF" {
            fmt.Println("用户退出了")
            return 0
        }
        fmt.Println("错误")
        return -1
    }
    return 1
}

func say(tcpConn *net.TCPConn) {
    for {
        data := make([]byte, 128)
        total, err := tcpConn.Read(data)
        fmt.Println(string(data[:total]), err)
        flag := checkErr(err)
        if flag == 0 {
            break
        }

        for _, conn := range ConnMap {
            if conn.RemoteAddr().String() == tcpConn.RemoteAddr().String() {
                continue
            }
            conn.Write(data[:total])
        }
    }
}

func main(){

    tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
    ConnMap = make(map[string]*net.TCPConn)
    for {
        tcpConn, _ := tcpListener.AcceptTCP()
        defer tcpConn.Close()
    
        ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
        fmt.Println("连接的客服端信息:", tcpConn.RemoteAddr().String())
        go say(tcpConn)
    }
}
