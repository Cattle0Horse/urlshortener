package main

import (
	"github.com/Cattle0Horse/url-shortener/application"
)

func main() {
	configFilePath := "./config/config.yaml"
	app, err := application.NewApp(configFilePath)
	if err != nil {
		panic(err)
	}
	app.Start()
}
