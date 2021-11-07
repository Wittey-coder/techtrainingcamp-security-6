package service

import (
	"awesomeProject/dao"
	"awesomeProject/tool"
)

// ClearOutdatedData 清理过期验证码服务
func ClearOutdatedData(expireTime int64) {
	userDao := dao.UserDao{Orm: tool.DbEngine}
	userDao.CleanOutdatedSmsCode(expireTime)
}
