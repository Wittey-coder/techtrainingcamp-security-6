package controller

import (
	"awesomeProject/service"
	"awesomeProject/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct {
}

func SendCode(context *gin.Context) {
	var applyCode = tool.ApplyCodeRequest{}
	err := context.ShouldBindJSON(&applyCode)
	if err != nil {
		context.JSON(http.StatusOK, tool.ApplyCodeResponse{
			Code:    1,
			Message: "信息错误！",
			Data: tool.Data{
				VerifyCode:   "",
				ExpireTime:   0,
				DecisionType: 0,
			},
		})
		return
	}

	smsService := service.SmsService{}
	hasSent, s := smsService.SendCode(applyCode.PhoneNumber)
	if hasSent {
		context.JSON(http.StatusOK, tool.ApplyCodeResponse{
			Code:    0,
			Message: "请求成功",
			Data: tool.Data{
				VerifyCode:   s,
				ExpireTime:   1000,
				DecisionType: 0,
			},
		})
		return
	}
	context.JSON(http.StatusOK, tool.ApplyCodeResponse{
		Code:    1,
		Message: "生成验证码失败！",
		Data: tool.Data{
			VerifyCode:   "",
			ExpireTime:   0,
			DecisionType: 0,
		},
	})
}
