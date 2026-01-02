package database

import (
	"server/global"
	"server/model/appTypes"

	"github.com/google/uuid"
)

type User struct {
	global.MODEL
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`
	Username  string            `json:"username"`
	Password  string            `json:"password"`
	Email     string            `json:"email"`
	Openid    string            `json:"openid"`
	Avatar    string            `json:"avatar"`
	Address   string            `json:"address"`
	Signature string            `json:"signature" gorm:"default:'签名是空白的，这位用户似乎比较低调。'"`
	RoleID    appTypes.RoleID   `json:"role"`
	Register  appTypes.Register `json:"register"`
	IsActive  bool              `json:"is_active"`
}
