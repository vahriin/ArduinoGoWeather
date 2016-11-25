package arduino

import (
	"fmt"
	"github.com/schleibinger/sio"
	"bufio"
)



type PortConnect struct {
	port *sio.Port
	reader *bufio.Reader
}

func Open(file string) (pc PortConnect, err error) {
	pc.port, err = sio.Open(file, 0x1002) //Baud Rate 115200
	if err != nil {
		panic(err)
	}
	pc.reader = bufio.NewReader(pc.port)
	return pc, nil
}

func (pc *PortConnect) ReadData() (string, error) {
	answer, err := pc.reader.ReadBytes('@')
	if err != nil {
		fmt.Print(err)
	}
	return string(answer), err
}

func (pc *PortConnect) Close() {
	pc.port.Close()
}

