package arduino

import (
	"github.com/schleibinger/sio"
	"bufio"
	"log"
)



type PortConnect struct {
	port *sio.Port
	reader *bufio.Reader
}

func Open(file string) (pc PortConnect, err error) {
	pc.port, err = sio.Open(file, 0x1002) //Baud Rate 115200
	if err != nil {
		return pc, err
	}
	pc.reader = bufio.NewReader(pc.port)
	//log.Println("arduino::Open: Open port ", pc.port)
	return pc, nil
}

func (pc *PortConnect) Gets() (string, error) {
		answer, err := pc.reader.ReadBytes('@')
	if err != nil {
		log.Print(err)
	}
	//log.Println("arduino: Get message: ", answer)
	return string(answer), err
}

func (pc *PortConnect) Close() {
	pc.port.Close()
}

