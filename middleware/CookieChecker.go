package middleware

import (
	"awesomeProject/session"
	"github.com/gin-gonic/gin"
)

func CookieChecker(context *gin.Context) {
	// 获取cookie里的sessionId数据
	sessionId, err := context.Cookie("sessionID")
	if err == nil {
		sessionInstance, err := session.Manager.Get(sessionId)
		if err != nil {
			return
		}
		// 验证此Session是否登录，用于登出时的验证
		isLogin, _ := (*sessionInstance).Get("isLogin")
		context.Set("isLogin", isLogin)
		// 验证此Session的用户名，用于注销时的验证
		username, _ := (*sessionInstance).Get("username")
		context.Set("username", username)
		session.Manager.Delete(sessionId)
	}
}
