package main

import (
	"log"
	"mediasoft/lesson8/internal/rest/app"
	"mediasoft/lesson8/internal/rest/config"
)

func main() {
	if err := app.Run(config.NewConfig()); err != nil {
		log.Println("error app.Run(): ", err)
	}
}
