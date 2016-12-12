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
	if err != nil {
		panic(err)
	}
	for {
		weath, err := spi.Gets()
		log.Println("collector::LoopGet: ", weath)
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
			log.Println("collector: Put message to Selector")
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
			for {
				select {
				case <-serverChannel:
					continue
				default:
					break Sel
				}
			}
		case serverChannel<-currentWeather:
			continue
		}
	}
}