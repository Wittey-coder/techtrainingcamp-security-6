/*
 * @Author: flwfdd
 * @Date: 2021-11-04 23:00:09
 * @LastEditTime: 2021-11-05 22:43:13
 * @Description:
 * _(:з」∠)_
 */
package risk

import (
	"awesomeProject/parameter"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const DEBUG = true

const (
	SUCCESS_CODE = 0
	FAILED_CODE  = 1
	CHECK_CODE   = 2
)

const (
	ENTER = 0
	CHECK = 1
	ABORT = 2
)

const (
	MAX_HEAT = 23332333
)

type heat_type struct {
	heat int //热力值
	time int //时间
}

//存放热力值
var mp map[string]heat_type

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

func get_time() int {
	return int(time.Now().Unix())
}

//获取热力值
func get_heat(s *string) (heat, dt int) {
	elem, ok := mp[*s]
	if ok {
		dt := get_time() - elem.time + 1
		if DEBUG {
			fmt.Printf("%v %v\n", *s, max(elem.heat-dt, 0))
		}
		return max(elem.heat-dt, 0), dt
	} else {
		elem = heat_type{0, get_time()}
		mp[*s] = elem
		return 0, 1
	}
}

//设置热力值
func set_heat(s *string, heat int) {
	mp[*s] = heat_type{min(max(heat, 0), MAX_HEAT), get_time()}
}

func CheckRisk(context *gin.Context) {
	value, _ := context.Get("JSON")

	status := check_risk_once(value)
	switch status {
	case ENTER:
		return
	case CHECK:
		context.Set("RETURN", parameter.ApplyCodeResponse{
			Code:    CHECK_CODE,
			Message: "需要进行验证！",
		})
		context.Abort()
	case ABORT:
		context.Set("RETURN", parameter.ApplyCodeResponse{
			Code:    FAILED_CODE,
			Message: "操作被拦截！",
		})
		context.Abort()
	}
}

//检查风险，返回0:正常 1:验证 2:拦截
func check_risk_once(i interface{}) int {
	s := ""

	switch info := i.(type) {
	case parameter.Environment: //IP和设备ID
		s = "I" + info.IP
		heat1, dt := get_heat(&s)
		set_heat(&s, 2*heat1/dt+4)
		s = "D" + info.DeviceID
		heat2, dt := get_heat(&s)
		set_heat(&s, 2*heat2/dt+4)
		heat := max(heat1, heat2)

		if heat < 300 {
			return 0
		} else if heat < 3600 {
			return 1
		} else {
			return 2
		}
	case parameter.LoginByPasswordRequest: //密码登录
		s = "U" + info.Username
		heat, dt := get_heat(&s)
		set_heat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 300 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else {
			out = ABORT
		}
		return max(out, check_risk_once(info.Environment))
	case parameter.LoginByPhoneRequest: //手机登录
		s = "P" + info.PhoneNumber
		heat, dt := get_heat(&s)
		set_heat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 300 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else {
			out = ABORT
		}
		return max(out, check_risk_once(info.Environment))
	case parameter.ApplyCodeRequest: //手机号获取验证码
		s = "P" + info.PhoneNumber
		heat, dt := get_heat(&s)
		set_heat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 300 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else {
			out = ABORT
		}
		return max(out, check_risk_once(info.Environment))
	case parameter.RegisterRequest: //手机号注册
		s = "P" + info.PhoneNumber
		heat, dt := get_heat(&s)
		set_heat(&s, 5*heat/dt+30)

		out := ENTER
		if heat < 300 {
			out = ENTER
		} else if heat < 3600 {
			out = CHECK
		} else {
			out = ABORT
		}
		return max(out, check_risk_once(info.Environment))
	default:
		fmt.Printf("Risk unkown type: %T\n", info)
		return 2
	}
}

func init() {
	mp = make(map[string]heat_type)
	fmt.Println("Here is RiskManagement")
}
