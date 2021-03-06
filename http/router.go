package http

import (
	"os"
	"static_service/http/controller"

	"github.com/gin-gonic/gin"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//routes
	router.POST("healthy", controller.Healthy)
	//TODO: 考虑单独迁出成为static服务
	statics := router.Group("static")
	{
		staticRootVALS := os.Getenv("STATIC_ROOT") + "/val"
		staticRootValIMG := os.Getenv("STATIC_ROOT") + "/img"
		//TODO:添加次数统计中间件
		statics.Static("/val/", staticRootVALS)
		statics.Static("/img/", staticRootValIMG)
	}
	upload := router.Group("upload")
	{
		upload.POST("img", controller.UploadImg)
	}
	return
}
