package cyclic

import (
	"awesomeProject/service"
	"awesomeProject/session"
	"github.com/robfig/cron"
	"log"
)

// AsynchronousLoop 设置异步循环任务
func AsynchronousLoop() {
	c := cron.New()
	err := c.AddFunc("* * * * * *", func() {
		// 定时为每秒执行一次
		// 所有的异步循环回调服务都丢这个地方
		service.ClearOutdatedData(100) //为了调试方便可以设置短一点的时间
	})
	err = c.AddFunc("* * * * * *", func() {
		err := session.Manager.Save()
		if err != nil {
			log.Fatal(err.Error())
		}
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	c.Start()
}
