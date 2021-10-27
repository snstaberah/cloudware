package model

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"inspur.com/cloudware/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库链接单例
var DB *gorm.DB

//定义自己的Writer
type MyLogger struct {
	mlog *logrus.Logger
}

//实现gorm/logger.Writer接口
func (m *MyLogger) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

func NewMyLogger() *MyLogger {
	log := util.LogrusLogger(os.Getenv("LOG_LEVEL"))
	return &MyLogger{mlog: log}
}

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	// 初始化GORM日志配置
	newLogger := logger.New(
		// log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		NewMyLogger(),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	// Error
	if connString == "" || err != nil {
		util.Log().Error("sqlite lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("sqlite lost: %v", err)
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	DB = db

	migration()
}
