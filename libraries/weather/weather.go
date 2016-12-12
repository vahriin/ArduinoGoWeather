package weather

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

func Create(data string) (Weather, error) {
	var weather Weather
	err := weather.Parse(data)
	return weather, err
}

func (weather *Weather)Parse(data string) (error) {
	var err error
	if (data == "") {
		return errors.New("No data string")
	}
	dataArray := strings.Split(data[:len(data)-1], "&")
	if len(dataArray) != 4 {
		return errors.New("error of data parsing")
	}

	weather.temperature, err = strconv.Atoi(dataArray[0])
	if err != nil {
		return err
	}
	weather.humidity, err = strconv.Atoi(dataArray[1])
	if err != nil {
		return err
	}
	press, err := strconv.Atoi(dataArray[2])
	if err != nil {
		return err
	}
	weather.pressure = uint(press) * 10
	weather.rain, err = rainConvert(dataArray[3])
	if err != nil {
		return err
	}
	return nil
}


func (weather Weather) TempCelsius() (float64) {
	return float64(weather.temperature)/10.0
}

func (weather Weather) TempFahrenheit() (float64) {
	return 9*(float64(weather.temperature)/10.0)/5+32
}

func (weather Weather) HumPercent() (float64) {
	return float64(weather.humidity)/10.0
}

func (weather Weather) PresPascal() (float64) {
	return float64(weather.pressure)
}

func (weather Weather) PresMmHg() (float64) {
	return float64(weather.pressure)/133.3224
}

func (weather Weather) PresBar() (float64) {
	return float64(weather.pressure)/100
}

func (weather Weather) RainType() (string) {
	return weather.rain
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
