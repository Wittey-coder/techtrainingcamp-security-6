package main

import (
	"awesomeProject/router"
	"awesomeProject/tool"
	"log"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := router.SetRoute()
	err = app.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}
