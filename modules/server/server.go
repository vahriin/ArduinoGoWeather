package server

import (
	"github.com/vahriin/Aroofino/modules/manager"
	"github.com/vahriin/Aroofino/libraries/weather"
	"log"
)

const (
	protocol = "tcp4"
)

func Server (weatherData <-chan weather.Weather, port string) {
	listener, err := manager.WaitConnect(protocol, port)
	if err != nil {
		log.Println("server: ", err)
	}
	manager.ConnectManager(listener, weatherData)
}