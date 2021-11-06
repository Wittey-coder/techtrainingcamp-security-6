package router

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.ReturnJSON, gin.Recovery(), gin.Logger())
	applyCode := engine.Group("/code") // 发送验证码
	{
		applyCode.POST("", controller.SendCode)
	}

	register := engine.Group("/register") // 注册相关的功能
	{
		register.POST("", controller.Register)
	}

	login := engine.Group("/login") // 登录相关的功能
	{
		login.POST("/sms", controller.LoginByPhone)
		login.POST("/word", controller.LoginByPassword)
	}

	cancel := engine.Group("/cancel") // 登出相关的功能
	{
		cancel.POST("/logout", controller.Logout)
		cancel.POST("/logoff", controller.Logoff)
	}

	return engine
}
