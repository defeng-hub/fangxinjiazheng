package v1

import (
	"fangxinjiazheng/models"
	"github.com/gin-gonic/gin"
)

type GlobalSettingController struct{}

func (GlobalSettingController) GetIndexAd(c *gin.Context) {
	var indexad []models.IndexAd
	models.DB.Find(&indexad)
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "获取首页广告",
		Data:  indexad,
	})
}

func (GlobalSettingController) AddIndexAd(c *gin.Context) {
	var indexobj models.IndexAd

	err := c.BindJSON(&indexobj)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "请求参数有误",
			Data:  nil,
		})
		return
	}
	tx := models.DB.Create(&indexobj)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "创建轮播图失败",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "创建轮播图成功",
		Data:  nil,
	})
	return
}

func (GlobalSettingController) DelIndexAd(c *gin.Context) {
	req := struct {
		ID uint `json:"id"`
	}{}
	err := c.BindJSON(&req)
	if err != nil || req.ID == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "请求参数错误",
			Data:  nil,
		})
		return
	}
	tx := models.DB.Delete(&models.IndexAd{ID: req.ID})
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "广告图不存在",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "删除成功",
		Data:  nil,
	})
	return

}

func (GlobalSettingController) ChangeIndexAd(c *gin.Context) {
	indexAd := models.IndexAd{}

	err := c.BindJSON(&indexAd)
	if err != nil || indexAd.ID == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "请求参数错误",
			Data:  nil,
		})
		return
	}

	tx := models.DB.First(&models.IndexAd{ID: indexAd.ID})
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "修改失败,请检查轮播图是否存在",
			Data:  nil,
		})
		return
	}
	models.DB.Save(&indexAd)
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "修改成功",
		Data:  nil,
	})
	return
}
