package router

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {
	engine := gin.Default()
	register := engine.Group("/register")
	{
		register.POST("/code", controller.SendCode)
	}

	return engine
}
