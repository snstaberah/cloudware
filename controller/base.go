package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
	"inspur.com/cloudware/conf"
	"inspur.com/cloudware/serializer"
	"inspur.com/cloudware/util"
)

type Sheldon struct {
	Name   string `json:"name" binding:"required,min=2,max=30"`
	Salary int    `json:"salary"`
	Sex    bool   `json:"sex"`
}

// Ping 状态检查接口
// @Summary Ping
// @Description 状态检查接口
// @Accept  json
// @Produce  json
// @param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200
// @Failure 500
// @Router /api/v1/ping [post]
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// TestLog 测试日志接口
// @Summary TestLog
// @Description 测试日志接口
// @Accept  json
// @Produce  json
// @param name path string true "Sheldon's salary is 80w"
// @Success 200
// @Failure 500
// @Router /api/v1/testlog [get]
func TestLog(c *gin.Context) {
	testgetparam := c.Query("name")
	if testgetparam != "" {
		util.Log().Error("%s's salary is 80w", testgetparam)
	}
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Sheldon's salary is 80w",
	})
}

// TestPost 测试post接口  不定义类型   直接取参数
// @Summary TestPost
// @Description 测试post接口
// @Accept  json
// @Produce  json
// @param name path string true "Sheldon's salary is 80w"
// @Success 200
// @Failure 500
// @Router /api/v1/TestPost [post]
func TestPost(c *gin.Context) {
	user := c.DefaultPostForm("player", "Sheldon")
	salary := c.PostForm("salary")

	util.Log().Error("%s's salary is 80w", user)
	util.Log().Error("%s's salary is %s", user, salary)
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Sheldon's salary is 80w",
	})
}

// TestPostStruct 测试post接口  不定义类型   直接取参数
// @Summary TestPostStruct
// @Description 测试post接口
// @Accept  json
// @Produce  json
// @param name path string true "Sheldon's salary is 80w"
// @Success 200
// @Failure 500
// @Router /api/v1/TestPostStruct [post]
func TestPostStruct(c *gin.Context) {
	p := Sheldon{}
	//要求body格式为json   自动解析为代码中的对象
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"msg":  "出错！",
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "success",
			"data": p,
		})
	}
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
