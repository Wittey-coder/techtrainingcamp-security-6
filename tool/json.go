package tool

type Environment struct {
	IP string `json:"ip"`
	DeviceID string `json:"device_id"`
}

type Data struct {
	VerifyCode string `json:"verify_code"`
	ExpireTime int `json:"expire_time"`
	DecisionType int `json:"decision_type"`
}

type ApplyCodeRequest struct {
	Environment `json:"environment"`
	PhoneNumber string `json:"phone_number"`
}

type ApplyCodeResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data `json:"data"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	VerifyCode string `json:"verify_code"`
	Environment `json:"environment"`
}

type RegisterResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	SessionId string `json:"session_id"`
	Data `json:"data"`
}

type LoginByPasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Environment `json:"environment"`
}

type LoginByPhoneRequest struct {
	PhoneNumber string `json:"phone_number"`
	VerifyCode string `json:"verify_code"`
	Environment `json:"environment"`
}

type LoginResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	SessionId string `json:"session_id"`
	Data `json:"data"`
}

type LogoutRequest struct {
	SessionId string `json:"session_id"`
	ActionType int `json:"action_type"`
	Environment `json:"environment"`
}

type LogoutResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}