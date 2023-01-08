package routers

import (
	v1 "fangxinjiazheng/controller/v1"
	"fangxinjiazheng/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	routers := r.Group("/user")
	{
		routers.POST("/admin/login", v1.UserController{}.LoginByAdmin)
		routers.GET("/test", middlewares.JWT(), v1.UserController{}.Logined)
		routers.POST("/loginByWx", v1.UserController{}.LoginByWx)
	}
}

func EmployeeController(r *gin.Engine) {
	routers := r.Group("/employee")
	{
		routers.POST("/addEmployee", middlewares.JWT(), v1.EmployeeController{}.AddEmployee)
		routers.POST("/getAllEmployee", middlewares.JWT(), v1.EmployeeController{}.GetAllEmployee)
		routers.POST("/delEmployee", middlewares.JWT(), v1.EmployeeController{}.DelEmployee)
		routers.POST("/changeEmployee", middlewares.JWT(), v1.EmployeeController{}.ChangeEmployee)
		routers.POST("/exportExcel", middlewares.JWT(), v1.EmployeeController{}.ExportExcel)
	}

}
