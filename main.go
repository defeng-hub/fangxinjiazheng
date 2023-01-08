package main

import (
	"fangxinjiazheng/middlewares"
	_ "fangxinjiazheng/models"
	"fangxinjiazheng/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
)

type Article struct {
	Title string
	Desc  string
}

func main() {
	r := gin.Default()
	//r := gin.New()

	//跨域
	r.Use(middlewares.CrosHandler())
	r.Static("/static", "static")

	//注册路由
	routers.InitRoutersInit(r)
	routers.UserRoutersInit(r)
	routers.ProductRoutersInit(r)
	routers.EmployeeController(r)
	routers.OrderRoutersInit(r)

	//读取conf.ini配置
	config, cerr := ini.Load("./conf/app.ini")
	fmt.Println("应用名:", config.Section("").Key("app_name").String())
	if cerr != nil {
		fmt.Printf("配置读取错误:%v", cerr)
		os.Exit(1)
	}

	//配置启动口
	err := r.Run(":8099")
	if err != nil {
		panic("程序启动失败")
	}
}
