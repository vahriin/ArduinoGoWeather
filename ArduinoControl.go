package main

import (
	"fmt"
	"github.com/schleibinger/sio"
	//"syscall"
	"bufio"
)

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

func (pc *PortConnect) setSensor(dhtExist bool, dhtType uint8, dhtPin uint8,
							     bmpExist bool, bmpType uint8, bmpPin uint8,
								 rainExist bool, rainPin uint8) {

	arguments := make([]byte, 8)
	arguments[0], arguments[3], arguments[6] = boolToByte(dhtExist), boolToByte(dhtExist), boolToByte(rainExist)
	arguments[1], arguments[2] = byte(dhtType), byte(dhtPin)
	arguments[4], arguments[5] = byte(bmpType), byte(bmpPin)
	arguments[7] = byte(rainPin)

	//var value int
	value, err := pc.port.Write(arguments)
	sayPanic(err)
	if value != 8{
		fmt.Println("less argument")
	}
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
	a.beginConnect("/dev/ttyACM0", 0x1002) //Baud Rate 115200
	a.setSensor(true, 22, 6, true, 28, 10, true, 0)
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

func boolToByte(q bool) (byte){
	if q {
		return byte(1)
	}else{
		return byte(0)
	}
}