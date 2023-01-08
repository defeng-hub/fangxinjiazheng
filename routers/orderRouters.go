package routers

import (
	v1 "fangxinjiazheng/controller/v1"
	"fangxinjiazheng/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutersInit(r *gin.Engine) {
	routers := r.Group("/orders")
	{
		routers.GET("/getAllOrders", middlewares.JWT(), v1.OrderController{}.GetAllOrder)
		routers.GET("/getAllOrdersByUser", middlewares.JWT(), v1.OrderController{}.GetAllOrdersByUser)
		routers.POST("/addOrder", middlewares.JWT(), v1.OrderController{}.AddOrder)
		routers.POST("/getOrderById", middlewares.JWT(), v1.OrderController{}.GetOrderById)
	}
}
