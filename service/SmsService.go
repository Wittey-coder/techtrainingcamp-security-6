package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/tool"
	"fmt"
	"math/rand"
	"time"
)

type SmsService struct {

}

func (s *SmsService) SendCode(phone string) (bool, string) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	smsCode := model.SmsCode{Phone: phone, Code: code, CreateTime: time.Now().Unix()}
	userDao := dao.UserDao{Orm: tool.DbEngine}
	result := userDao.InsertCode(smsCode)
	return result > 0, phone
}
