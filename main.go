package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"inspur.com/cloudware/conf"
	"inspur.com/cloudware/server"
	"inspur.com/cloudware/util"
)

func initLog() (err error) {

	// 设置fmt实现的日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 设置logrus日志级别
	logrusLogger := util.LogrusLogger(os.Getenv("LOG_LEVEL"))

	// gin支持release  debug  test, release模式，打印信息少
	gin_mode := os.Getenv("GIN_MODE")
	gin.SetMode(gin_mode)                // 默认debug模式，打印信息多
	gin.DefaultWriter = logrusLogger.Out // gin框架自己记录的日志使用logrus输出到文件   默认输出到控制台
	return
}

// @title 这里写标题`
// @version 1.0`
// @description 这里写描述信息`
// @termsOfService [http://swagger.io/terms/](http://swagger.io/terms/)`
// @contact.name 这里写联系人信息`
// @contact.url [http://www.swagger.io/support](http://www.swagger.io/support)`
// @contact.email support@swagger.io`
// @license.name Apache 2.0`
// @license.url [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`
// @host 这里写接口服务的host`
// @BasePath 这里写base path`
func main() {

	// 从配置文件读取配置
	conf.Init()
	//初始化日志
	err := initLog()
	if err != nil {
		panic(err)
	}
	// 装载路由
	r := server.NewRouter()
	port := os.Getenv("SERVER_PORT")
	r.Run(port)
}
