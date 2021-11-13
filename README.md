# techtrainingcamp-security-6

A secure login machine

## Backgound

字节后端训练营-抓到你了-第六组项目大作业

## Environment

- Windows or MacOS or Linux
- Go(1.13+)
- MySQL
- Redis

## Badge

![](https://img.shields.io/badge/redigo-1.6.3-red)  ![](https://img.shields.io/badge/gin-1.7.4-blue)  ![](https://img.shields.io/badge/mysql-1.6.0-green)  ![](https://img.shields.io/badge/xorm-0.7.9-red)  ![](https://img.shields.io/badge/cron-1.2.0-yellow)  ![](https://img.shields.io/badge/go.uuid-1.2.0-red)

## Install

进入项目根目录，需要编译，输入：

```shell
$ go build main.go
```

## Usage

编译完成后，会生成main文件，需要运行服务器时，输入：

```shell
$ ./main
```

## Configuration

需要按照自己的需要配置app.json

```json
{
  "app_name": "VerifySecurity",
  "app_mode": "debug或者release",
  "app_host": "主机IP",
  "app_port": "占用端口，默认8080",
  "session_buffer": "redis或者memory，memory模式Session没有过期时间",
  "database": {
    "driver": "使用的数据库",
    "user": "root",
    "password": "密码",
    "host": "主机IP",
    "port": "占用端口",
    "db_name": "数据库名",
    "charset": "字符集",
    "show_sql":bool值，是否在终端显示SQL
  },
  "redis": {
    "password": "密码",
    "host": "主机IP",
    "port": "端口"
  }
}
```

示例：

```json
{
  "app_name": "VerifySecurity",
  "app_mode": "debug",
  "app_host": "localhost",
  "app_port": "8080",
  "session_buffer": "redis",
  "database": {
    "driver": "mysql",
    "user": "root",
    "password": "1314",
    "host": "127.0.0.1",
    "port": "3306",
    "db_name": "ginsql",
    "charset": "utf8",
    "show_sql":false
  },
  "redis": {
    "password": "123",
    "host": "127.0.0.1",
    "port": "6379"
  }
}
```



## Interface

- 手机验证码请求                RESTful接口：POST /code
- 注册请求                            RESTful接口：POST /register
- 手机登录请求                    RESTful接口：POST /login/sms
- 密码登录请求                    RESTful接口：POST /login/word
- 登出/注销请求                   RESTful接口：POST /logout

```go
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
	DecisionType int `json:"decision_type"`
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
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    `json:"data"`
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
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    `json:"data"`
}

// LogoutRequest 登出/注销请求
type LogoutRequest struct {
	ActionType  int `json:"action_type"`
	Environment `json:"environment"`
}

// LogoutResponse 登出/注销响应
type LogoutResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

```



## Maintainers
@[**Wittey-dev**](https://github.com/Wittey-dev)

@[**flwfdd**](https://github.com/flwfdd)



