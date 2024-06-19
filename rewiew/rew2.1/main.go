package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=WatherGetterer
type WatherGetterer interface {
	GetData(s string) (*http.Response, error)
	GetWather(sity string) string
}

// https://goweather.herokuapp.com/weather/Moscow
type MyClient struct {
	client        *http.Client
	watherService WatherGetterer
}

func (c *MyClient) GetData(s string) (*http.Response, error) {
	resp, err := c.client.Get(s)
	return resp, err
}

func (c *MyClient) GetWather(sity string) string {
	resp, err := c.GetData("https://goweather.herokuapp.com/weather/" + sity)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	oData := &Whetherer{}
	_ = json.NewDecoder(resp.Body).Decode(oData)
	return fmt.Sprintf("Temperature todey: %s, tomorow: %s, next day: %s, next next: %s", oData.Temperature, oData.Forecast[0].Temperature,
		oData.Forecast[1].Temperature, oData.Forecast[2].Temperature)
}

type Whetherer struct {
	Temperature string     `json:"temperature"`
	Wind        string     `json:"wind"`
	Description string     `json:"description"`
	Forecast    []Forecast `json:"forecast"`
}

type Forecast struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}

func main() {
	c := &http.Client{}
	m := &MyClient{
		client: c,
	}
	fmt.Println(m.GetWather("Moscow"))
}
