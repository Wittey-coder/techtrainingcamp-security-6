package session

import (
	"errors"
	"fmt"
)

var Manager SessionManager

// Init 初始化外部变量
func Init(provider string, addr string, options ...string) error {
	// 选择版本，选择redis版，内存版没有过期时间功能
	switch provider {
	case "memory":
		Manager = NewMemorySessionManager()
	case "redis":
		Manager = NewRedisSessionManager()
	default:
		fmt.Errorf("not supported")
		return errors.New("not supported")
	}
	err := Manager.Init(addr, options...)
	return err
}
