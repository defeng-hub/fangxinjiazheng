package routers

import (
	v1 "fangxinjiazheng/controller/v1"
	"fangxinjiazheng/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductRoutersInit(r *gin.Engine) {
	routers := r.Group("/product")
	{
		routers.GET("/allProduct", v1.ProductController{}.GetAllProduct)
		routers.POST("/getProductById", v1.ProductController{}.GetProductById)
		routers.POST("/getProductByName", v1.ProductController{}.GetProductByName)

		routers.POST("/delProduct", middlewares.JWT(), v1.ProductController{}.DelProductById)
		routers.POST("/addProduct", middlewares.JWT(), v1.ProductController{}.AddProduct)
		routers.POST("/changeProduct", middlewares.JWT(), v1.ProductController{}.ChangeProduct)
	}
}
