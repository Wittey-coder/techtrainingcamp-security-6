package dao

import (
	"awesomeProject/model"
	"awesomeProject/tool"
	"log"
)

type UserDao struct {
	*tool.Orm
}

func (u *UserDao) InsertCode(sms model.SmsCode) int64 {
	result, err := u.InsertOne(&sms)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
