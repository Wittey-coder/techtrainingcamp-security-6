package controller

import (
	"awesomeProject/parameter"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS       int = 0
	FAILED        int = 1
	ACTION_LOGOUT     = 1
	ACTION_LOGOFF     = 2
)

// SendCode 发送验证码
func SendCode(context *gin.Context) {
	var applyCode = parameter.ApplyCodeRequest{}
	err := context.ShouldBindJSON(&applyCode)
	if err != nil {
		context.JSON(http.StatusOK, parameter.ApplyCodeResponse{
			Code:    FAILED,
			Message: "信息错误！",
			CodeData: parameter.CodeData{
				VerifyCode:   "",
				ExpireTime:   0,
				DecisionType: 0,
			},
		})
		return
	}

	smsService := service.UserService{}
	hasSent, s := smsService.SendCode(applyCode.PhoneNumber)
	if hasSent {
		context.JSON(http.StatusOK, parameter.ApplyCodeResponse{
			Code:    SUCCESS,
			Message: "请求成功",
			CodeData: parameter.CodeData{
				VerifyCode:   s,
				ExpireTime:   1000,
				DecisionType: 0,
			},
		})
		return
	}
	context.JSON(http.StatusOK, parameter.ApplyCodeResponse{
		Code:    FAILED,
		Message: "生成验证码失败！",
		CodeData: parameter.CodeData{
			VerifyCode:   "",
			ExpireTime:   0,
			DecisionType: 0,
		},
	})
}

// LoginByPhone 用手机验证码登录
func LoginByPhone(context *gin.Context) {
	var loginByPhoneParameter = parameter.LoginByPhoneRequest{}
	err := context.ShouldBindJSON(&loginByPhoneParameter)
	if err != nil {
		context.JSON(http.StatusOK, parameter.LoginResponse{
			Code:      FAILED,
			Message:   "发送的表单格式错误",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
	}
	smsService := service.UserService{}
	user := smsService.LoginByPhone(loginByPhoneParameter)
	if user != nil {
		context.JSON(http.StatusOK, parameter.LoginResponse{
			Code:      SUCCESS,
			Message:   "登录成功",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		return
	}
	context.JSON(http.StatusOK, parameter.LoginResponse{
		Code:      FAILED,
		Message:   "登录失败",
		SessionId: "",               // TODO
		Data:      parameter.Data{}, // TODO
	})
}

// LoginByPassword 用密码登录
func LoginByPassword(context *gin.Context) {
	// TODO
}

// Register 注册
func Register(context *gin.Context) {
	// TODO
}

// Logout 登出
func Logout(context *gin.Context) {
	// TODO
}

// Logoff 注销
func Logoff(context *gin.Context) {
	// TODO
}
