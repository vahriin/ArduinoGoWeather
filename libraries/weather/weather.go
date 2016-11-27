package main

import(
	"strings"
	"errors"
	"strconv"
)

const (
	SHOWER = 400
	RAIN = 800
	DRIZZLE = 1000
	DRY = 1024
)

type Weather struct {
	temperature int //temperature and humidity values are stores with
	humidity int    //one characters after the decimal point
	pressure uint   //pressure value is stores in the normal form
	rain string		//rain value stores a force of rain
}

func (weather Weather) Fahrenheit() (int) {
	return 9*weather.temperature/5+320
}

func (weather Weather) MmHg() (float64) {
	return float64(weather.pressure)/133.3224
}

func (weather Weather) Bar() (float64) {
	return float64(weather.pressure)/100
}

func Parse(data string) (Weather, error) {
	var weather Weather;
	var err error
	if (data == "") {
		return weather, errors.New("No data string")
	}
	dataArray := strings.Split(data[:len(data)-1], "&")
	if len(dataArray) != 4 {
		return weather, errors.New("error of data parsing")
	}

	weather.temperature, err = strconv.Atoi(dataArray[0])
	if err != nil {
		return weather, err
	}
	weather.humidity, err = strconv.Atoi(dataArray[1])
	if err != nil {
		return weather, err
	}
	temp, err := strconv.Atoi(dataArray[2])
	if err != nil {
		return weather, err
	}
	weather.pressure = uint(temp) * 10
	weather.rain, err = rainConvert(dataArray[3])
	if err != nil {
		return weather, err
	}
	return weather, nil
}

func (weather Weather) GetData(temperature, humidity, pressure, rain bool, formats ...int) (string, error) {
	//use bool vars for get the corresponding values.
	//formats var includes format for getting values:
	//temperature: 0 for celsius, 1 for fahrenheit
	//pressure: 0 for pascals, 1 for mmHg, 2 for Bar
	//default all formats is 0
	//if boolean var is false, not set corresponding values.
	//set formats in the order of the corresponding boolean vars

	var returnData string

	if temperature {
		var temp int
		if len(formats) != 0 {
			temp = formats[0]
		} else {
			temp = 0
		}
		switch temp {
		case 0 :
			returnData += strconv.FormatFloat(float64(weather.temperature)/10.0, 'f', -1, 32)
		case 1 :
			returnData += strconv.FormatFloat(float64(weather.Fahrenheit())/10.0, 'f', -1, 32)
		}
		returnData += " "
	}

	if humidity {
		returnData += strconv.FormatFloat(float64(weather.humidity)/10.0, 'f', -1, 32)
		returnData += " "
	}

	if pressure {
		var temp int //index of parametrs in formats
		if len(formats) == 1 && !temperature {
			temp = formats[0]
		} else if len(formats) == 2 && temperature {
			temp = formats[1]
		} else {
			temp = 0
		}
		switch temp {
		case 0:
			returnData += strconv.FormatFloat(float64(weather.pressure), 'f', -1, 32)
		case 1:
			returnData += strconv.FormatFloat(weather.MmHg(), 'f', -1, 32)
		case 2:
			returnData += strconv.FormatFloat(weather.Bar(), 'f', -1, 32)
		}
		returnData += " "
	}

	if rain {
		returnData += weather.rain
	}
	return returnData, nil
}

func rainConvert(rain string) (string, error) {
	irain, err := strconv.Atoi(rain)
	if err != nil {
		return "", err
	}

	if irain <= 0 || irain > 1023 {
		return "", errors.New("Error of the rainSensor")
	} else if irain < SHOWER {
		return "Shower", nil
	} else if irain < RAIN {
		return "Rain", nil
	} else if irain < DRIZZLE {
		return  "Drizzle", nil
	} else if irain < DRY {
		return "Dry", nil
	}
	return "", errors.New("Unexpected error")
}
