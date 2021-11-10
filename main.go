package main

import (
	"awesomeProject/cyclic"
	"awesomeProject/router"
	"awesomeProject/session"
	"awesomeProject/tool"
	"log"
)

func main() {
	// 解析app.json
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	// 配置数据库
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// 配置路由
	app := router.SetRoute()

	// 设置定时任务
	cyclic.AsynchronousLoop()

	// 配置Session管理器
	err = session.Init(cfg.SessionBuffer, cfg.Redis.Host+":"+cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	// GO!
	err = app.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
