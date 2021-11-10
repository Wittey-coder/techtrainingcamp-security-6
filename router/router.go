package router

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"awesomeProject/risk"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {
	engine := gin.New()
	/*
		middleware.ReturnJSON 用于统一返回Json
		middleware.CookieChecker 用于解析客户端携带的Cookie
	*/
	engine.Use(middleware.ReturnJSON, middleware.CookieChecker, gin.Recovery(), gin.Logger())
	/*
		controller.GetSendCodeJSON 解析客户端携带的Json
		risk.CheckRiskForApplyCode 以频率为基准的风控系统
		controller.SendCode 发送验证码的控制器
	*/
	applyCode := engine.Group("/code") // 发送验证码
	{

		applyCode.POST("", controller.GetSendCodeJSON, risk.CheckRiskForApplyCode, controller.SendCode)
	}
	/*
		controller.GetRegisterJSON 解析客户端携带的Json
		middleware.SessionController 用于统一生成Session
		risk.CheckRiskForApplyCode 以频率为基准的风控系统
		controller.Register 管理注册的控制器
	*/
	register := engine.Group("/register") // 注册相关的功能
	{
		register.POST("", controller.GetRegisterJSON, middleware.SessionController, risk.CheckRiskForLog, controller.Register)
	}
	/*
		controller.GetLoginByPhoneJSON 解析客户端携带的Json
		controller.GetLoginByPasswordJSON 解析客户端携带的Json
		middleware.SessionController 用于统一生成Session
		risk.CheckRiskForApplyCode 以频率为基准的风控系统
		controller.LoginByPhone 管理登录的控制器
		controller.LoginByPassword 管理登录的控制器
	*/
	login := engine.Group("/login") // 登录相关的功能
	{
		login.POST("/sms", controller.GetLoginByPhoneJSON, middleware.SessionController, risk.CheckRiskForLog, controller.LoginByPhone)
		login.POST("/word", controller.GetLoginByPasswordJSON, middleware.SessionController, risk.CheckRiskForLog, controller.LoginByPassword)
	}

	/*
		controller.GetLogoutJSON 解析客户端携带的Json
		middleware.SessionController 用于统一生成Session
		controller.Logout 管理登出的控制器
	*/
	cancel := engine.Group("/logout") // 登出相关的功能
	{
		cancel.POST("/", controller.GetLogoutJSON, middleware.SessionController, controller.Logout)
	}

	return engine
}
