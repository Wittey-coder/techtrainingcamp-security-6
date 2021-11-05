package parameter

// Environment 设备信息
type Environment struct {
	IP       string `json:"ip"`
	DeviceID string `json:"device_id"`
}

// CodeData 验证码及相关信息
type CodeData struct {
	VerifyCode   string `json:"verify_code"`
	ExpireTime   int    `json:"expire_time"`
	DecisionType int    `json:"decision_type"`
}

// Data Session及相关信息
type Data struct {
	SessionId    string `json:"session_id"`
	ExpireTime   int    `json:"expire_time"`
	DecisionType int    `json:"decision_type"`
}

// ApplyCodeRequest 验证码请求
type ApplyCodeRequest struct {
	Environment `json:"environment"`
	PhoneNumber string `json:"phone_number"`
}

// ApplyCodeResponse 验证码响应
type ApplyCodeResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeData `json:"code_data"`
}

// RegisterRequest  注册请求
type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	VerifyCode  string `json:"verify_code"`
	Environment `json:"environment"`
}

// RegisterResponse 注册响应
type RegisterResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	SessionId string `json:"session_id"`
	Data      `json:"data"`
}

// LoginByPasswordRequest 密码登录请求
type LoginByPasswordRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Environment `json:"environment"`
}

// LoginByPhoneRequest  手机登录请求
type LoginByPhoneRequest struct {
	PhoneNumber string `json:"phone_number"`
	VerifyCode  string `json:"verify_code"`
	Environment `json:"environment"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	SessionId string `json:"session_id"`
	Data      `json:"data"`
}

// LogoutRequest 登出/注销请求
type LogoutRequest struct {
	SessionId   string `json:"session_id"`
	ActionType  int    `json:"action_type"`
	Environment `json:"environment"`
}

// LogoutResponse 登出/注销响应
type LogoutResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
