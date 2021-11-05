package model

type SmsCode struct {
	Id         int64  `xorm:"pk autoincr" json:"id"`     // 自增ID
	Phone      string `xorm:"varchar(11)" json:"phone"`  // 手机号
	Code       string `xorm:"varchar(6)" json:"code"`    // 验证码
	CreateTime int64  `xorm:"bigint" json:"create_time"` // 生成时间，用来清除过期验证码
}
