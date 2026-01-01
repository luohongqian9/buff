package database

import "server/global"

type Login struct {
	global.MODEL
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	LoginMethod string `json:"login_method"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	OS          string `json:"os"`
	BrowserInfo string `json:"browser_info"`
	DeviceInfo  string `json:"device_info"`
	status      int    `json:"status"`
}
