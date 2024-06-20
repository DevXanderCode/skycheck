package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=5edb5652cbae4f958cc115437242006&q=Port-harcourt&aqi=no&days=1&alerts=no")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather Api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
