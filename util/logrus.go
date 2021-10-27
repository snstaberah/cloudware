package util

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func LogrusLogger(level string) *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出同时到控制台和文件
	mw := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(mw)

	//设置日志级别
	intLevel := logrus.DebugLevel
	switch level {
	case "error":
		intLevel = logrus.ErrorLevel
	case "warning":
		intLevel = logrus.WarnLevel
	case "info":
		intLevel = logrus.InfoLevel
	case "debug":
		intLevel = logrus.DebugLevel
	}

	logger.SetLevel(intLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
