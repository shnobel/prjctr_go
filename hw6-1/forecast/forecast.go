package forecast

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type ConstantForecast struct {
	City string
}

func (cf ConstantForecast) GetTommorowsWeather() (string, error) {
	return fmt.Sprintf("%s: ☀️   +72°F", cf.City), nil
}

type RandomForecast struct {
	City        string
	temperature int
}

func (rf *RandomForecast) GetTommorowsWeather() (string, error) {
	rf.temperature = rand.Intn(15-5) + 5
	return fmt.Sprintf("%s: ☀️   +%d°F", rf.City, rf.temperature), nil
}

type WetherAPIForecast struct {
	City string
}

func (wf WetherAPIForecast) GetTommorowsWeather() (string, error) {
	url := fmt.Sprintf("https://wttr.in/%s?format=3", wf.City)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
