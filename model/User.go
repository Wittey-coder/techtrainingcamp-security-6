package model

type User struct {
	UUID         string `xorm:"pk" json:"uuid"`                             // uuid
	Username     string `xorm:"varchar(20) notnull unique" json:"username"` // 用户名
	Phone        string `xorm:"varchar(11) notnull unique" json:"phone"`    // 电话号码，11位
	Password     string `xorm:"varchar(255) notnull" json:"password"`       // 密码
	RegisterTime int64  `xorm:"bigint created" json:"register_time"`        // 注册时间，默认为插入时间
	ActiveNumber int    `xorm:"tinyint default 0" json:"active_number"`     // 当前已登录设备数量
}
