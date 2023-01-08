package controller

import (
	"fangxinjiazheng/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"path"
	"strconv"
	"time"
)

type InitController struct{}

func (InitController) InitDB(c *gin.Context) {
	if err := models.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Img{},
		&models.Employee{}, &models.IndexAd{}); err != nil {
		c.String(200, "迁移失败")
		return
	}
	c.String(200, "数据库迁移成功, ")

}

func (InitController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, models.Rep{
			State:   -1,
			Msg:     "请求数据非法",
			Data:    nil,
			Content: nil,
		})
		return
	}

	//获取文件的拓展名
	extName := path.Ext(file.Filename)

	//被允许的后缀map
	allowExtMap := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		c.JSON(200, models.Rep{
			State:   -2,
			Msg:     "数据非法",
			Data:    nil,
			Content: nil,
		})

		return
	}
	today := time.Now().Format("20060102")

	//创建多层夹
	err2 := os.MkdirAll("./static/upload/"+today, os.ModePerm)
	if err2 != nil {
		c.JSON(200, gin.H{
			"code": -3,
			"msg":  "保存失败",
		})
		return
	}

	//要保存到的地址
	filepath := path.Join("./static/upload/"+today, strconv.Itoa(int(time.Now().Unix()))+"-"+file.Filename)
	err1 := c.SaveUploadedFile(file, filepath)
	if err1 != nil {
		c.JSON(200, models.Rep{
			State:   -3,
			Msg:     "保存失败",
			Data:    nil,
			Content: nil,
		})
		return
	}
	config, _ := ini.Load("./conf/app.ini")
	api := config.Section("").Key("api").String()
	res := "http://" + api + "/" + filepath
	//保存到数据库
	tx := models.DB.Save(&models.Img{Url: "/" + filepath})
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State:   -1,
			Msg:     "保存失败",
			Data:    nil,
			Content: nil,
		})
		return
	}
	m := map[string]string{}
	m["publicUrl"] = res
	m["localUrl"] = "/" + filepath
	c.JSON(200, models.Rep{
		State:   0,
		Msg:     "保存成功",
		Data:    m,
		Content: nil,
	})
	return
}

// GetTurnover 获取最近10天营业额
func (InitController) GetTurnover(c *gin.Context) {
	var res []interface{}
	// 0代表今天  1代表1天前
	for i := 0; i < 10; i++ {
		var orders []models.Order
		sqls := fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(created_at) = %d", i)
		models.DB.Where(sqls).Find(&orders)
		m := make(map[string]interface{})
		m["title"] = fmt.Sprintf("%d天前", i)
		if m["title"] == "0天前" {
			m["title"] = "今天"
		} else if m["title"] == "1天前" {
			m["title"] = "昨天"
		}
		var zong float64
		for _, orderobj := range orders {
			if orderobj.PayStatus != 0 {
				zong += orderobj.Amount
			}
		}
		m["data"] = zong
		res = append(res, m)
	}

	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "近10天营业额获取成功",
		Data:  res,
	})
}
