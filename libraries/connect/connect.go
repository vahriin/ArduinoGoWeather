package connect

import(
	"net"
	"bufio"
	"time"
)

const (
	deadline = 30
)

type ClientConnection struct { //data client and buffers for work
	conn *net.TCPConn
	reader bufio.Reader
	writer bufio.Writer
}

func StartConnection(connect *net.TCPConn) (ClientConnection) {
	var cc ClientConnection
	cc.conn = connect
	cc.reader = *bufio.NewReader(cc.conn)
	cc.writer = *bufio.NewWriter(cc.conn)
	cc.conn.SetDeadline(time.Now().Add(time.Minute*deadline))
	return cc
}

func (cc ClientConnection) Gets() (string, error) {
	answer, _,  err := cc.reader.ReadLine()
	if err != nil {
		return "", err
	}
	cc.conn.SetDeadline(time.Now().Add(time.Minute*deadline))
	return string(answer), nil
}

func (cc ClientConnection) Puts(message string) (error) {
	_, err := cc.writer.WriteString(message)
	cc.writer.Flush()
	cc.conn.SetDeadline(time.Now().Add(time.Minute*deadline))
	return  err
}

func (cc ClientConnection) Close() {
	cc.conn.Close()
}






