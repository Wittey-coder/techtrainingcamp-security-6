package dao

import (
	"awesomeProject/model"
	"awesomeProject/tool"
	"log"
	"time"
)

type UserDao struct {
	*tool.Orm
}

// InsertCode 添加验证码到服务器
func (u *UserDao) InsertCode(sms model.SmsCode) int64 {
	result, err := u.InsertOne(&sms)
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

// ValidateSmsCode 查看是否有此手机验证码数据对在数据库中
func (u *UserDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	_, err := u.Where("phone = ?, code = ?", phone, code).Get(&sms)
	if err != nil {
		log.Println(err.Error())
	}
	return &sms
}

// QueryByPhone 查看此手机是否被注册
func (u *UserDao) QueryByPhone(phone string) *model.User {
	var user model.User
	has, err := u.Where("phone = ?", phone).Get(&user)
	if err != nil {
		log.Println(err.Error())
	}
	if has {
		return &user
	} else {
		return nil
	}
}

// InsertUser 添加用户
func (u *UserDao) InsertUser(user model.User) int64 {
	result, err := u.InsertOne(&user)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return result
}

// CleanOutdatedSmsCode
func (u *UserDao) CleanOutdatedSmsCode(expireTime int64) {
	_, err := u.Where("create_time < ? + ?", time.Now().Unix(), expireTime).Delete(model.SmsCode{})
	if err != nil {
		log.Println(err.Error())
	}
}
