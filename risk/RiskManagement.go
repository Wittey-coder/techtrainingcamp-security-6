package risk

import (
	"awesomeProject/parameter"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const DEBUG = true

const (
	ENTER = iota
	CHECK
	WAIT
	BANNED
)

const (
	SUCCESS = iota
	FAILED
)

const (
	MAX_HEAT = 23332333
)

type heatType struct {
	heat int //热力值
	time int //时间
}

//存放热力值
var mp map[string]heatType

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getTime() int {
	return int(time.Now().Unix())
}

//获取热力值
func getHeat(s *string) (heat, dt int) {
	elem, ok := mp[*s]
	if ok {
		dt := getTime() - elem.time + 1
		if DEBUG {
			fmt.Printf("%v %v\n", *s, max(elem.heat-dt, 0))
		}
		return max(elem.heat-dt, 0), dt
	} else {
		elem = heatType{0, getTime()}
		mp[*s] = elem
		return 0, 1
	}
}

//设置热力值
func setHeat(s *string, heat int) {
	mp[*s] = heatType{min(max(heat, 0), MAX_HEAT), getTime()}
}

func CheckRiskForApplyCode(context *gin.Context) {
	value, _ := context.Get("JSON")

	status := checkRiskOnce(value)
	switch status {
	case ENTER:
		return
	case CHECK:
		context.Set("RETURN", parameter.ApplyCodeResponse{
			Code:    FAILED,
			Message: "需要进行验证！",
			CodeData: parameter.CodeData{
				VerifyCode:   "",
				ExpireTime:   0,
				DecisionType: CHECK,
			},
		})
		context.Set("decisionType", CHECK)
		context.Abort()
	case WAIT:
		context.Set("RETURN", parameter.ApplyCodeResponse{
			Code:    FAILED,
			Message: "访问频率过高！需要等待热力值下降！",
			CodeData: parameter.CodeData{
				VerifyCode:   "",
				ExpireTime:   0,
				DecisionType: WAIT,
			},
		})
		context.Set("decisionType", WAIT)
		context.Abort()
	case BANNED:
		context.Set("RETURN", parameter.ApplyCodeResponse{
			Code:    FAILED,
			Message: "操作被拦截！",
			CodeData: parameter.CodeData{
				VerifyCode:   "",
				ExpireTime:   0,
				DecisionType: BANNED,
			},
		})
		context.Set("decisionType", BANNED)
		context.Abort()
	}
}

func CheckRiskForLog(context *gin.Context) {
	value, _ := context.Get("JSON")

	status := checkRiskOnce(value)
	switch status {
	case ENTER:
		return
	case CHECK:
		context.Set("RETURN", parameter.LoginResponse{
			Code:    0,
			Message: "需要进行验证!",
			Data: parameter.Data{
				DecisionType: CHECK,
			},
		})
		context.Set("decisionType", CHECK)
		context.Abort()
	case WAIT:
		context.Set("RETURN", parameter.LoginResponse{
			Code:    FAILED,
			Message: "访问频率过高！需要等待热力值下降！",
			Data: parameter.Data{
				DecisionType: WAIT,
			},
		})
		context.Set("decisionType", WAIT)
		context.Abort()
	case BANNED:
		context.Set("RETURN", parameter.LoginResponse{
			Code:    FAILED,
			Message: "操作被拦截！",
			Data: parameter.Data{
				DecisionType: BANNED,
			},
		})
		context.Set("decisionType", BANNED)
		context.Abort()
	}
}

//检查风险，返回 0:正常 1:验证 2:等待 2:拦截
func checkRiskOnce(i interface{}) int {
	s := ""

	switch info := i.(type) {
	case parameter.Environment: //IP和设备ID
		s = "I:" + info.IP
		heat1, dt := getHeat(&s)
		setHeat(&s, 2*heat1/dt+4)
		s = "D" + info.DeviceID
		heat2, dt := getHeat(&s)
		setHeat(&s, 2*heat2/dt+4)
		heat := max(heat1, heat2)

		out := ENTER
		if heat < 50 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else if heat < 12000 {
			out = WAIT
		} else {
			out = BANNED
		}
		return out
	case parameter.LoginByPasswordRequest: //密码登录
		s = "U" + info.Username
		heat, dt := getHeat(&s)
		setHeat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 50 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else if heat < 12000 {
			out = WAIT
		} else {
			out = BANNED
		}
		return max(out, checkRiskOnce(info.Environment))
	case parameter.LoginByPhoneRequest: //手机登录
		s = "P" + info.PhoneNumber
		heat, dt := getHeat(&s)
		setHeat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 50 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else if heat < 12000 {
			out = WAIT
		} else {
			out = BANNED
		}
		return max(out, checkRiskOnce(info.Environment))
	case parameter.ApplyCodeRequest: //手机号获取验证码
		s = "P" + info.PhoneNumber
		heat, dt := getHeat(&s)
		setHeat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 50 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else if heat < 12000 {
			out = WAIT
		} else {
			out = BANNED
		}
		return max(out, checkRiskOnce(info.Environment))
	case parameter.RegisterRequest: //手机号注册
		s = "P" + info.PhoneNumber
		heat, dt := getHeat(&s)
		setHeat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 50 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else if heat < 12000 {
			out = WAIT
		} else {
			out = BANNED
		}
		return max(out, checkRiskOnce(info.Environment))
	default:
		fmt.Printf("Risk unkown type: %T\n", info)
		return 2
	}
}

func init() {
	mp = make(map[string]heatType)
	fmt.Println("Here is RiskManagement")
}
