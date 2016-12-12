package manager

import(
	"net"
	"github.com/vahriin/Aroofino/libraries/reqresp"
	"github.com/vahriin/Aroofino/libraries/connect"
	"github.com/vahriin/Aroofino/libraries/weather"
	"log"
)

func WaitConnect(protocol, port string) (*net.TCPListener, error) { //this is Server main function
	address, err := net.ResolveTCPAddr(protocol, ":" + port) //create listen address
	if err != nil {
		return nil, err
	}
	openedPort, err := net.ListenTCP(protocol, address) //get listener port
	if err != nil {
		return nil, err
	}
	log.Println("manager: Listen port ", port)
	log.Println("manager: Wait for connect")
	return openedPort, nil
}

func ConnectManager(listenPort *net.TCPListener, weatherChannel <-chan weather.Weather) {
	for{
		connection, err := listenPort.AcceptTCP()//accept connection from client
		if err != nil {
			continue
		}
		log.Println("manager: Set a connect")
		cc := connect.StartConnection(connection)
		go ClientOperate(cc, weatherChannel) //work with client
	}
}

func ClientOperate(client connect.ClientConnection, weatherChannel <-chan weather.Weather) {
	defer client.Close()
	for{
		request, err := client.Gets()
		if err != nil {
			log.Println(err)
			client.Close()
			return
		}
		log.Println("manager: Get Request: ", request)
		response := reqresp.MakeResponse(request, <-weatherChannel)
		err = client.Puts(response)
		if err != nil {
			log.Println(err)
			client.Close()
			return
		}
		log.Println("manager: Put Response: ", response)
	}
}