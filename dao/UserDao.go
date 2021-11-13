package dao

import (
	"awesomeProject/model"
	"awesomeProject/tool"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

type UserDao struct {
	*tool.Orm
}

// InsertCode 添加验证码到服务器
func (u *UserDao) InsertCode(sms *model.SmsCode) int64 {
	result, err := u.InsertOne(sms)
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

// ValidateSmsCode 查看是否有此手机验证码数据对在数据库中
func (u *UserDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	_, err := u.Where("phone = ? AND code = ?", phone, code).Get(&sms)
	if err != nil {
		log.Println(err.Error())
	}
	return &sms
}

// ValidatePassword 查看用户名和密码是否正确
func (u *UserDao) ValidatePassword(username string, password string) *model.User {
	var user model.User
	has, err := u.Where("username = ? AND password = ?", username, password).Get(&user)
	if err != nil {
		log.Println(err.Error())
	}
	if has {
		return &user
	} else {
		return nil
	}
}

//func (u *UserDao) UpdateLoggedDeviceNumber(changedDevices int, username string) *model.User {
//	var user model.User
//	_, _ = u.Where("username = ?", username).Get(&user)
//	var updated model.User
//	updated.ActiveNumber = user.ActiveNumber + changedDevices
//	_, err := u.Id(user.Id).Cols("active_number").Update(updated)
//	if err != nil {
//		return nil
//	}
//	return &user
//}

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
func (u *UserDao) InsertUser(username string, phone string, password string) (int64, *model.User) {

	var user = model.User{
		Id:       uuid.NewV4().String(),
		Username: username,
		Phone:    phone,
		Password: password,
	}

	result, err := u.InsertOne(&user)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}
	return result, &user
}

func (u *UserDao) DeleteUserByUsername(username string) error {
	_, err := u.Where("username = ?", username).Delete(model.User{})
	if err != nil {
		return err
	}
	return nil
}

// CleanOutdatedSmsCode 清除过期数据
func (u *UserDao) CleanOutdatedSmsCode(expireTime int64) {
	_, err := u.Where("create_time < ? - ?", time.Now().Unix(), expireTime).Delete(model.SmsCode{})
	if err != nil {
		log.Println(err.Error())
	}
}

func (u *UserDao) QueryCodeByPhone(phone string) *model.SmsCode {
	var code model.SmsCode
	has, err := u.Where("phone = ?", phone).Get(&code)
	if err != nil {
		log.Println(err.Error())
	}
	if has {
		return &code
	} else {
		return nil
	}
}
