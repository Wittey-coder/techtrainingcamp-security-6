package service

import (
	"awesomeProject/dao"
	"awesomeProject/tool"
)

type ClearOutdatedService struct {
}

// ClearData 清理过期验证码服务
func (c *ClearOutdatedService) ClearData(expireTime int64) {
	userDao := dao.UserDao{Orm: tool.DbEngine}
	userDao.CleanOutdatedSmsCode(expireTime)
}
