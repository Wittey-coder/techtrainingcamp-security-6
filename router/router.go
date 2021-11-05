package router

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {
	engine := gin.Default()

	applyCode := engine.Group("/code") // 发送验证码
	{
		applyCode.GET("", controller.SendCode)
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
