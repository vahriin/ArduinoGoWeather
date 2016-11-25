package main

import (
	"github.com/vahriin/Aroofino/libraries/arduino"
	"fmt"
)

func main() {
	a, err := arduino.Open("/dev/ttyACM0")
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 10; i++ {
		str, _ := a.ReadData()
		fmt.Println(str)
	}

}
