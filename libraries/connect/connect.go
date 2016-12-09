package connect

import(
	"net"
	"bufio"
)

const buffer = 1000

type ClientConnection struct { //data about client and buffers for work
	number int
	conn *net.TCPConn
	reader bufio.Reader
	writer bufio.Writer
}

func (cc ClientConnection) Gets() (string, error) {
	answer, _,  err := cc.reader.ReadLine()
	if err != nil {
		return "", err
	}
	return string(answer), nil
}

func (cc ClientConnection) Puts(message string) (error) {
	_, err := cc.writer.WriteString(message)
	return  err
}

//func (cc ClientConnection)




