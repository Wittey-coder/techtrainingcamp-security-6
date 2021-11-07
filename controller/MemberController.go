package controller

import (
	"awesomeProject/parameter"
	"awesomeProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	SUCCESS int = 0
	FAILED  int = 1
)

// GetSendCodeJSON 解析发送验证码JSON
func GetSendCodeJSON(context *gin.Context) {
	var applyCode = parameter.ApplyCodeRequest{}
	err := context.ShouldBindWith(&applyCode, binding.JSON)
	if err != nil {
		context.JSON(http.StatusOK, parameter.ApplyCodeResponse{
			Code:    FAILED,
			Message: "信息错误！",
			CodeData: parameter.CodeData{
				VerifyCode:   "",
				ExpireTime:   600,
				DecisionType: 0,
			},
		})
		context.Abort()
	} else {
		context.Set("JSON", applyCode)
	}
}

// SendCode 发送验证码
func SendCode(context *gin.Context) {
	applyCode_, _ := context.Get("JSON")
	applyCode := applyCode_.(parameter.ApplyCodeRequest)

	smsService := service.UserService{}
	hasSent, s := smsService.SendCode(applyCode.PhoneNumber)
	if hasSent {
		context.Set("RETURN", parameter.ApplyCodeResponse{
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
	context.Set("RETURN", parameter.ApplyCodeResponse{
		Code:    FAILED,
		Message: "生成验证码失败！",
		CodeData: parameter.CodeData{
			VerifyCode:   "",
			ExpireTime:   0,
			DecisionType: 0,
		},
	})
}

// GetLoginByPhoneJSON 解析手机登录JSON
func GetLoginByPhoneJSON(context *gin.Context) {
	var loginByPhoneParameter = parameter.LoginByPhoneRequest{}
	err := context.ShouldBindWith(&loginByPhoneParameter, binding.JSON)
	if err != nil {
		context.Set("RETURN", parameter.LoginResponse{
			Code:      FAILED,
			Message:   "发送的表单格式错误",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		context.Abort()
	} else {
		context.Set("JSON", loginByPhoneParameter)
	}
}

// LoginByPhone 用手机验证码登录
func LoginByPhone(context *gin.Context) {
	loginByPhoneParameter_, _ := context.Get("JSON")
	loginByPhoneParameter := loginByPhoneParameter_.(parameter.LoginByPhoneRequest)

	smsService := service.UserService{}
	user := smsService.LoginByPhone(loginByPhoneParameter)
	if user != nil {
		context.Set("RETURN", parameter.LoginResponse{
			Code:      SUCCESS,
			Message:   "登录成功",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		return
	}
	context.Set("RETURN", parameter.LoginResponse{
		Code:      FAILED,
		Message:   "登录失败",
		SessionId: "",               // TODO
		Data:      parameter.Data{}, // TODO
	})
}

// GetLoginByPasswordJSON 解析密码登录JSON
func GetLoginByPasswordJSON(context *gin.Context) {
	var loginByPasswordParameter = parameter.LoginByPasswordRequest{}
	err := context.ShouldBindWith(&loginByPasswordParameter, binding.JSON)
	if err != nil {
		context.Set("RETURN", parameter.LoginResponse{
			Code:      FAILED,
			Message:   "发送的表单格式错误",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		context.Abort()
	} else {
		context.Set("JSON", loginByPasswordParameter)
	}
}

// LoginByPassword 用密码登录
func LoginByPassword(context *gin.Context) {
	loginByPasswordParameter_, _ := context.Get("JSON")
	loginByPasswordParameter := loginByPasswordParameter_.(parameter.LoginByPasswordRequest)

	userService := service.UserService{}
	user := userService.LoginByPassword(loginByPasswordParameter)
	if user != nil {
		context.Set("RETURN", parameter.LoginResponse{
			Code:      SUCCESS,
			Message:   "登录成功",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		return
	}
	context.Set("RETURN", parameter.LoginResponse{
		Code:      FAILED,
		Message:   "登录失败",
		SessionId: "",               // TODO
		Data:      parameter.Data{}, // TODO
	})
}

// GetRegisterJSON 解析注册JSON
func GetRegisterJSON(context *gin.Context) {
	var registerParameter = parameter.RegisterRequest{}
	err := context.ShouldBindWith(&registerParameter, binding.JSON)
	if err != nil {
		context.Set("RETURN", parameter.LoginResponse{
			Code:      FAILED,
			Message:   "发送的表单格式错误",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		context.Abort()
	} else {
		context.Set("JSON", registerParameter)
	}
}

// Register 注册
func Register(context *gin.Context) {
	registerParameter_, _ := context.Get("JSON")
	registerParameter := registerParameter_.(parameter.RegisterRequest)

	userService := service.UserService{}
	register := userService.Register(registerParameter)
	if register != nil {
		context.Set("RETURN", parameter.LoginResponse{
			Code:      SUCCESS,
			Message:   "注册成功",
			SessionId: "",               // TODO
			Data:      parameter.Data{}, // TODO
		})
		return
	}
	context.Set("RETURN", parameter.LoginResponse{
		Code:      FAILED,
		Message:   "注册失败，有相同的用户名或手机号被注册或者验证码错误！",
		SessionId: "",               // TODO
		Data:      parameter.Data{}, // TODO
	})
}

// Logout 登出
func Logout(context *gin.Context) {
	// TODO
}

// Logoff 注销
func Logoff(context *gin.Context) {
	// TODO
}
