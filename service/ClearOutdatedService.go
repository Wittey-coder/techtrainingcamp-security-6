package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/tool"
	"log"
	"time"
)

type ClearOutdatedService struct {
}

func (c *ClearOutdatedService) ClearData(expireTime int64) {
	userDao := dao.UserDao{Orm: tool.DbEngine}
	_, err := userDao.Where("create_time < ? + ?", time.Now().Unix(), expireTime).Delete(model.SmsCode{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
