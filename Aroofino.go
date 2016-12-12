package main

import (
	//"github.com/vahriin/Aroofino/libraries/arduino"
	"github.com/vahriin/Aroofino/modules/collector"
	"github.com/vahriin/Aroofino/libraries/weather"
	"github.com/vahriin/Aroofino/modules/manager"
	//"log"
	//"fmt"
	//"log"
)

func main() {
		listener, err := manager.WaitConnect("tcp4", "12064")
	if err != nil {
		panic(err)
	}

	dataCollector := make(chan weather.Weather)
	dataServer := make(chan weather.Weather, 10)
	dataSwap := make(chan string)

	go collector.LoopGet(dataSwap)
	go collector.MakeWeather(dataCollector, dataSwap)
	go collector.Selector(dataCollector, dataServer)
	go manager.ConnectManager(listener, dataServer)

	//var ss string
	//fmt.Scanln(&ss)
	for{}
}
