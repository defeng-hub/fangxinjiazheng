package routers

import (
	"fangxinjiazheng/controller"
	v1 "fangxinjiazheng/controller/v1"
	"fangxinjiazheng/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutersInit(r *gin.Engine) {
	r.POST("/upload", controller.InitController{}.Upload)
	r.POST("/database", controller.InitController{}.InitDB)
	r.POST("/getTurnover", controller.InitController{}.GetTurnover)
	r.GET("/getTurnover", controller.InitController{}.GetTurnover)
	r.POST("/system/getIndexAd", v1.GlobalSettingController{}.GetIndexAd)
	r.POST("/system/addIndexAd", middlewares.JWT(), v1.GlobalSettingController{}.AddIndexAd)
	r.POST("/system/delIndexAd", middlewares.JWT(), v1.GlobalSettingController{}.DelIndexAd)
	r.POST("/system/changeIndexAd", middlewares.JWT(), v1.GlobalSettingController{}.ChangeIndexAd)
}
