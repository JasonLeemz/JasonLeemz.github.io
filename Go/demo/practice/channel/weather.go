package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 京东万象
//const url = "https://way.jd.com/he/freeweather?city=beijing&appkey=24e88e485b4d9432867b0c487bea3efa"
const url = "https://way.jd.com/he/freeweather?city=%s&appkey=24e88e485b4d9432867b0c487bea3efa"

type WeatherResponse struct {
	Code   string `json:"code"`
	Charge bool   `json:"charge"`
	Msg    string `json:"msg"`
	Result Result `json:"result"`
}

type Update struct {
	Loc string `json:"loc"`
	Utc string `json:"utc"`
}
type Basic struct {
	City   string `json:"city"`
	Cnty   string `json:"cnty"`
	ID     string `json:"id"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
	Update Update `json:"update"`
}

type Cond struct {
	CodeD string `json:"code_d"`
	CodeN string `json:"code_n"`
	TxtD  string `json:"txt_d"`
	TxtN  string `json:"txt_n"`
}
type Tmp struct {
	Max string `json:"max"`
	Min string `json:"min"`
}
type Wind struct {
	Deg string `json:"deg"`
	Dir string `json:"dir"`
	Sc  string `json:"sc"`
	Spd string `json:"spd"`
}
type DailyForecast struct {
	Cond Cond   `json:"cond"`
	Date string `json:"date"`
	Hum  string `json:"hum"`
	Pcpn string `json:"pcpn"`
	Pop  string `json:"pop"`
	Pres string `json:"pres"`
	Tmp  Tmp    `json:"tmp"`
	Vis  string `json:"vis"`
	Wind Wind   `json:"wind"`
}
type HeWeather5 struct {
	Basic         Basic           `json:"basic"`
	DailyForecast []DailyForecast `json:"daily_forecast"`
}
type Result struct {
	HeWeather5 []HeWeather5 `json:"HeWeather5"`
}

func GetWeather(ctx context.Context, city string) (*WeatherResponse, error) {

	cityUrl := fmt.Sprintf(url, city)

	resp, err := http.Get(cityUrl)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	weather := &WeatherResponse{}

	err = json.Unmarshal(body, weather)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	//log.Println(weather.Result)
	return weather, nil
}
