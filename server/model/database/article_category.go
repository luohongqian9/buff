package database

import "server/global"

type ArticleCategory struct {
	global.MODEL
	Number int `json:"number"`
}
