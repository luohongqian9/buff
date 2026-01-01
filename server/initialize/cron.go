package initialize

import (
	"os"
	"server/global"
	"server/task"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keysAndValues", keysAndValues))
}
func (l *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Error(err), zap.Any("keysAndValues", keysAndValues))
}

func NewZaplogger() *ZapLogger {
	return &ZapLogger{
		logger: global.Log,
	}
}

func InitCron() {
	c := cron.New(cron.WithLogger(NewZaplogger()))
	if err := task.RegisterScheduledTasks(c); err != nil {
		global.Log.Error("定时任务注册失败", zap.Error(err))
		os.Exit(1)
	}
	c.Start()
}
