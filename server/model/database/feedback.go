package database

import "server/global"

type Feedback struct {
	global.MODEL
	UserUUID string `json:"user_uuid" gorm:"type:char(36)"`
	User     User   `json:"user" gorm:"foreignKey:UserUUID;references:UUID"`
	Content  string `json:"content"`
	Reply    string `json:"reply"`
}
