package database

import "server/global"

type ArticleTag struct {
	global.MODEL
	Number int `json:"number"`
}
