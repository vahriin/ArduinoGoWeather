package main

import (
	"net"

	"github.com/vahriin/Aroofino/libraries/connect"
	"log"
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
	log.Println("client: connection")
	server.Puts("1,1,1,1")
	log.Println(server.Gets())
}
