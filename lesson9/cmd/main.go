package main

import (
	"log"
	"mediasoft/lesson9/internal/app"
	"mediasoft/lesson9/internal/config"
)

func main() {
	if err := app.Run(config.NewConfig()); err != nil {
		log.Println("error app.Run(): ", err)
	}
}
