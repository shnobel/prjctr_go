package main

import (
	"fmt"
	"prjctr_go/hw6-1/forecast"
)

type WeatherForecaster interface {
	GetTommorowsWeather() (string, error)
}

func printForecast(wf WeatherForecaster) {
	forecast, _ := wf.GetTommorowsWeather()
	fmt.Println(forecast)
}

func main() {
	cf := forecast.ConstantForecast{
		City: "New York",
	}
	printForecast(cf)

	rf := forecast.RandomForecast{
		City: "New York",
	}
	printForecast(&rf)

	wf := forecast.WetherAPIForecast{
		City: "New York",
	}
	printForecast(wf)
}
