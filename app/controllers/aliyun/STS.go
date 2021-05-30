package aliyun

import (
	"fmt"
	"gin-aliyun/app/services"
	"github.com/gin-gonic/gin"
)

func GetAccess(c *gin.Context)  {
	res,err := services.GetAccess()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200,gin.H{
			"state":4001,
			"msg":"获取sts失败！",
		})
	}else {
		c.JSON(200, gin.H{
			"state":2000,
			"message":"success",
			"data": res,
		})
	}
}