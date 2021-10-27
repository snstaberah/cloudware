## Design by Sheldon，who's salary is 80w per year

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API,例如cloudware这种垃圾

## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](https://gorm.io/index.html): ORM工具。本项目需要配合sqlite使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. 自行实现了国际化i18n的一些基本功能
8. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证
9. [logrus](https://github.com/sirupsen/logrus) 集成流行的日志工具logrus，并设置gin和gorm都使用logrus打印，默认同时打印到控制台和文件中
10.[jwt-go](https://github.com/dgrijalva/jwt-go) 在老bei的强力要求下使用jwt-go实现bearer token认证

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/user/register```用户注册接口
3. 实现了```/api/v1/user/login```用户登录接口
4. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)
6. 实现了一些test接口，可以参考进行get/post等处理

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象，可以参考使用
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件
9. logs目录存放日志文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
SQLITE_DB="test.db"        #默认集成了sqlite   可以方便的切换为mysql  可通过路径设置db文件位置
REDIS_ADDR="127.0.0.1:6379"
REDIS_PW=""                #集成了  但暂时用不上
REDIS_DB=""
SESSION_SECRET="setOnProducation"
GIN_MODE="debug"           #gin框架的日志级别   debug打的多   release打的少
LOG_LEVEL="debug"          #logrus日志级别   debug打的多   
SERVER_PORT=":3000"
```



## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init cloudware
go mod tidy
外网环境可以使用
export GOPROXY=http://mirrors.aliyun.com/goproxy/
办公网环境可以使用代理
export GOPROXY=http://10.7.12.185:8089/repository/go-group/
研发内网环境可以使用
export GOPROXY=http://100.2.97.200:8081/repository/go-group/
go run main.go // 自动安装
```

## 运行
环境需安装了sqlite

```shell
进入cloudware项目根目录
go run main.go
项目运行后默认启动在3000端口（可以修改，参考gin文档)
```

```docker
编译完成镜像后可使用docker运行
支持通过env传递覆盖Godotenv中定义的默认参数
docker run -d  --net host --env SERVER_PORT=":3003" cloudware:v0.1
```


## 编译
使用docker file 两段式编译
使用安装好sqlite的基础镜像  基于debian:buster-slim 
进入工程根目录  执行
docker build -f ./Dockerfile . -t cloudware:v0.1


