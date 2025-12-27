package initialize

import (
	"server/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	mysqlCfg := global.Config.Mysql

	db, err := gorm.Open(mysql.Open(mysqlCfg.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlCfg.LogLevel()),
	})
	if err != nil {
		global.Log.Error("初始化数据库失败", zap.Error(err))
		panic(err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	sqlDb.SetMaxOpenConns(mysqlCfg.MaxOpenConns)

	return db
}
