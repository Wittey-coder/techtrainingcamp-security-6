package model

type OnlineDevice struct {
	SessionId string `xorm:"varchar(128) pk" json:"session_id"`
}
