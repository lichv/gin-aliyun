package routers

import (
	"gin-aliyun/app/controllers/aliyun"
	"gin-aliyun/app/middlewares"
	"gin-aliyun/utils"
	"gin-aliyun/utils/setting"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	iconPath := path.Join(setting.AppSetting.RootPath, "favicon.ico")
	if utils.IsExist(iconPath) {
		r.Use(favicon.New(iconPath))
	}

	r.Use(middlewares.Cors())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"state":   2000,
			"message": "success",
		})
	})

	apiv1 := r.Group("/api/aliyun/v1")
	apiv1.Use()
	{
		apiv1.GET("/aliyun/getConfig", aliyun.GetConfig)
		apiv1.Any("/sts/getAccess",aliyun.GetAccess)
	}

	return r
}
