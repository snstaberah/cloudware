package host

import (
	"github.com/gin-gonic/gin"
	"inspur.com/cloudware/controller"
	"inspur.com/cloudware/model"
	"inspur.com/cloudware/serializer"
)

// AddHost
func AddHost(c *gin.Context) {
	host := model.Host{}
	//post body映射代码对象
	if err := c.ShouldBind(&host); err == nil {
		err := host.AddHost()
		if err == nil {
			c.JSON(200, serializer.Response{
				Code: 0,
				Msg:  "okok",
			})
		} else {
			c.JSON(400, controller.ErrorResponse(err))
		}

	} else {
		c.JSON(400, controller.ErrorResponse(err))
	}
}

// DeleteHost
func DeleteHost(c *gin.Context) {
	c.JSON(400, serializer.Response{
		Code: 0,
		Msg:  "wuha",
	})
}

// UpdateHost
func UpdateHost(c *gin.Context) {
	c.JSON(400, serializer.Response{
		Code: 0,
		Msg:  "wuha",
	})
}

// GetHostByID
func GetHostByID(c *gin.Context) {
	if hostID := c.Query("hostid"); hostID != "" {
		host, err := model.GetHost(hostID)
		if err == nil {
			c.JSON(200, host)
		} else {
			c.JSON(400, serializer.Response{
				Code: 0,
				Msg:  "wuha",
			})
		}

	} else {
		c.JSON(400, serializer.Response{
			Code: 0,
			Msg:  "wuha",
		})
	}

}
