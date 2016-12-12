package collector

import (
	"github.com/vahriin/Aroofino/libraries/arduino"
	"log"
	"github.com/vahriin/Aroofino/libraries/weather"
	"time"
)

const (
	readPause = 10 //in second
)

func LoopGet(weatherStr chan<- string) {
	spi, err := arduino.Open("/dev/ttyACM0")
	//log.Println(spi.Gets())
	if err != nil {
		panic(err)
	}
	for {
		//log.Println("loopGet work")
		weath, err := spi.Gets()
		log.Println(weath)
		if err != nil {
			log.Println(err)
		}
		weatherStr <- weath
	}
}

func MakeWeather(collChannel chan<- weather.Weather, weatherStr <-chan string){
	timer := time.NewTimer(time.Second*readPause)
	var currentWeather weather.Weather
	null := make(chan string, 10)
	var w string
	for {
		select {
		case w = <- weatherStr:
			null <- w
		case <-timer.C:
			currentWeather.Parse(w)
			collChannel <- currentWeather
			//log.Println("collector: Put message to Selector")
			timer.Reset(time.Second*readPause)
		case <- null:
		}
	}
}

func Selector(weatChannel <-chan weather.Weather, serverChannel chan weather.Weather) {
	currentWeather := <-weatChannel
	for {
		Sel:
		select {
		case currentWeather = <-weatChannel:
			//log.Println("colletcor: Selector get message")
			for {
				select {
				case <-serverChannel:
					//log.Println("collector: null message")
					continue
				default:
					break Sel
				}
			}
		case serverChannel<-currentWeather:
			//log.Println("collector: Put message to server channel")
			continue
		}
	}
}


/*func Collector(serverChannel chan weather.Weather, tty string){
	ard, _ := arduino.Open(tty)

	var swap *string
	dataChannel := make(chan weather.Weather)
	go LoopGet(ard, swap)
	go MakeWeather(dataChannel, swap)
	go Selector(dataChannel, serverChannel)
}*/