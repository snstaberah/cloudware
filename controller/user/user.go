package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"inspur.com/cloudware/controller"
	"inspur.com/cloudware/model"
	"inspur.com/cloudware/serializer"
	"inspur.com/cloudware/service"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	//post body映射代码对象
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, controller.ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.JWTLogin(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, controller.ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": gin.H{"woshisei": "sheldon", "salary": 800000},
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		//类型断言
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
