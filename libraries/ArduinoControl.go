package main

import (
	"fmt"
	"github.com/schleibinger/sio"
	"bufio"
)

const BaudRate = 0x1002 //Baud Rate 115200

type PortConnect struct {
	port *sio.Port
	reader *bufio.Reader
}

func (pc *PortConnect) beginConnect(file string, br uint32) (err error) {
	pc.port, err = sio.Open(file, br)
	sayPanic(err)
	pc.reader = bufio.NewReader(pc.port)
	return nil
}

func (pc *PortConnect) readData() (string, error) {
	answer, err := pc.reader.ReadBytes('@')
	sayError(err)
	return string(answer), err
}

func (pc *PortConnect) endConnect() {
	pc.port.Close()
}



func main() {
	var a PortConnect
	a.beginConnect("/dev/ttyACM0", BaudRate)
	for i := 0; i < 10; i++ {
		str, _ := a.readData()
		fmt.Println(str)
	}

}

func sayError(err error){
	if err != nil {
		fmt.Print(err)
	}
}

func sayPanic(err error){
	if err != nil {
		panic(err)
	}
}