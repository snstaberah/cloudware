package conf

import (
	"os"

	"github.com/joho/godotenv"
	"inspur.com/cloudware/model"
	"inspur.com/cloudware/util"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量  默认根目录.env文件
	godotenv.Load()

	// 读取对应语言的翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("SQLITE_DB"))
	// cache.Redis()
}
