package main

import (
	"net"

	"github.com/vahriin/Aroofino/libraries/connect"
	"log"
	"time"
	"fmt"
)

func main(){
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:12064")
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		log.Println(err)
	}
	server := connect.StartConnection(conn)
	fmt.Println(Send(server, "0,0,0,0\n"))
	fmt.Println(Send(server, "0,0,0,1\n"))
	fmt.Println(Send(server, "0,0,0,T\n"))
	fmt.Println(Send(server, "0,0,1,0\n"))
	fmt.Println(Send(server, "0,0,M,0\n"))
	fmt.Println(Send(server, "0,0,B,0\n"))
	fmt.Println(Send(server, "0,0,P,0\n"))
	fmt.Println(Send(server, "0,1,0,0\n"))
	fmt.Println(Send(server, "1,0,0,0\n"))
	fmt.Println(Send(server, "C,0,0,0\n"))
	fmt.Println(Send(server, "F,0,0,0\n"))
	fmt.Println(Send(server, "C,1,1,1\n"))
	fmt.Println(Send(server, "F,1,B,1\n"))

}

func Send(conn connect.ClientConnection, request string) (string) {
	conn.Puts(request)
	answer, err := conn.Gets()
	time.Sleep(time.Second*2)
	if err != nil {
		log.Println(err)
	}
	return answer
}