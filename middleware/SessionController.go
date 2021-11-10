package middleware

import (
	"awesomeProject/session"
	"github.com/gin-gonic/gin"
)

func SessionController(c *gin.Context) {
	c.Next()
	// 下发一个新的Id
	s, id, err := session.Manager.CreateSession()
	if err != nil {
		return
	}
	c.SetCookie("sessionID", id, 60, "/", "localhost:8080", false, true)
	// 获取是否登录的参数
	isLogin, exists := c.Get("isLogin")
	if exists {
		err := s.Set("isLogin", isLogin)
		if err != nil {
			return
		}
		err = s.Save()
		if err != nil {
			return
		}
	}
	// 获取风控参数
	decisionType, exists := c.Get("decisionType")
	if exists {
		err := s.Set("decisionType", decisionType)
		if err != nil {
			return
		}
		err = s.Save()
		if err != nil {
			return
		}
	}
	// 获取用户名参数
	username, exists := c.Get("username")
	if exists {
		err := s.Set("username", username)
		if err != nil {
			return
		}
		err = s.Save()
		if err != nil {
			return
		}
	}

}
