package api

import (
	"strings"
	"errors"
)

type ClientMessage struct { //communication between Client Operate and ...
	message string
	number int
	active bool
}

type Request struct{ //request for Sender
	tempFormat int
	humFormat int
	presFormat int
	rainFormat int
}

type SenderData struct { //TODO: rename me, pls
	Request
	clientChannel chan string
}

func MakeRequest(request string) (Request, error) {
	parameters := strings.Split(request, ",")
	if len(parameters) != 4 {
		return nil, errors.New("wrong number of parameters. Expected 4 parameters")
	}

	var answer Request

	switch parameters[0] { //temperature format parse
	case "0":
		answer.tempFormat = 0 //don't take the temperature
	case "1":
		fallthrough
	case "c":
		fallthrough
	case "C":
		answer.tempFormat = 1 //temp in Celsius
	case "f":
		fallthrough
	case "F":
		answer.tempFormat = 2 //temp in Fahrenheit
	default:
		return nil, errors.New("Unexpected argument in position 0")
	}

	switch parameters[1] { //humidity format parse
	case "0":
		answer.humFormat = 0
	case "1":
		fallthrough
	case "P":
		fallthrough
	case "p":
		answer.humFormat = 1 //hum in Percent
	default:
		return nil, errors.New("Unexpected argument in position 1")
	}

	switch parameters[2] {
	case "0":
		answer.presFormat = 0
	case "1":
		fallthrough
	case "M":
		fallthrough
	case "m":
		answer.presFormat = 1 //pres in Mm.Hg
	case "B":
		fallthrough
	case "b":
		answer.presFormat = 2 //pres in Bar
	case "P":
		fallthrough
	case "p":
		answer.presFormat = 3 //pres in Pascal
	default:
		return nil, errors.New("Unexpected argument in position 2")
	}

	switch parameters[3] {
	case "0":
		answer.rainFormat = 0
	case "1":
		answer.rainFormat = 1
	default:
		return nil, errors.New("Unexpected argument in position 3")
	}
	return answer, nil
}


