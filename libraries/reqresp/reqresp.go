package reqresp

import (
	"strings"
	"github.com/vahriin/Aroofino/libraries/weather"
	"strconv"
)

/*type ClientMessage struct { //communication between Client Operate and ...
	text string
}

func MakeMessage(request string) (ClientMessage) {
	var cm ClientMessage
	cm.text = request
	return cm
}

func (cm *ClientMessage) SetText(request string) {
	cm.text = request
}

func (cm ClientMessage) GetText() (string) {
	return cm.text
}*/

func MakeResponse(cm string, weatherData weather.Weather) (string) {
	response := ""
	request := strings.Split(strings.ToUpper(strings.Replace(cm, " ", "", -1)), ",")

	if len(request) != 4 {
		return "Wrong number of parameters. Expected 4 parameters"
	}

	switch request[0] {
	case "0":
		break
	case "1":
		fallthrough
	case "C": //temp in Celsius Format
		response += strconv.FormatFloat(weatherData.TempCelsius(), 'f', 1, 64)
	case "F": //temp in Fahrenheit Format
		response += strconv.FormatFloat(weatherData.TempFahrenheit(), 'f', 1, 64)
	default:
		return "Unexpected argument in position 0"
	}
	response += ","

	switch request[1] { //switch used for functionality expansion in the future
	case "0":
		break
	case "1":
		fallthrough
	case "P": //humidity in Percent format
		response += strconv.FormatFloat(weatherData.HumPercent(), 'f', 1, 64)
	default:
		return "Unexpected argument in position 1"
	}
	response += ","

	switch request[2] {
	case "0":
		break
	case "1":
		fallthrough
	case "M":
		response += strconv.FormatFloat(weatherData.PresMmHg(), 'f', 1, 64)
	case "B":
		response += strconv.FormatFloat(weatherData.PresBar(), 'f', 1, 64)
	case "P":
		response += strconv.FormatFloat(weatherData.PresPascal(), 'f', 1, 64)
	default:
		return "Unexpected argument in position 2"
	}
	response += ","

	switch request[3] {
	case "0":
		break
	case "1":
		fallthrough
	case "T": //rain in type
		response += weatherData.RainType()
	default:
		return "Unexpected argument in position 3"
	}
	return response
}



