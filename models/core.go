package models

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

//自动执行
func init() {
	config, cerr := ini.Load("./conf/app.ini")

	if cerr != nil {
		fmt.Printf("配置读取错误:%v\n", cerr)
		os.Exit(1)
	}
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		//Logger: 可以在这里自定义日志模板
	})

	if err != nil {
		//fmt.Printf("创建数据库连接失败:%v\n", err)
		panic(fmt.Sprintf("创建数据库连接失败:%v\n", err))
	} else {
		//
		fmt.Printf("数据库连接成功\n")
	}
}
