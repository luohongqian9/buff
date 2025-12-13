package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/config"
	"server/global"
	"strings"
	"time"
)

func InitMySQL(m config.Mysql) error {
	dsn := buildDSN(m)
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(parseLogMode(m.LogMode)),
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.DB = db
	return nil
}

func buildDSN(c config.Mysql) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.Config,
	)
}

func parseLogMode(mode string) logger.LogLevel {
	switch strings.ToLower(mode) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}
