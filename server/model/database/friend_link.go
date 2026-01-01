package database

import "server/global"

type FriendLink struct {
	global.MODEL
	Logo        string `json:"logo"`
	Image       Image  `json:"-" gorm:"foreignKey:Logo;references:URL"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}
