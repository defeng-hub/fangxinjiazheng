package v1

import (
	"fangxinjiazheng/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (ProductController) GetAllProduct(c *gin.Context) {
	var pros []models.Product
	tx := models.DB.Find(&pros)
	if tx.Error != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "获取失败",
		})
		c.Abort()
		return
	}
	c.JSON(200, models.Rep{
		State:   0,
		Msg:     "获取成功",
		Data:    pros,
		Content: nil,
	})
}

func (ProductController) DelProductById(c *gin.Context) {
	reqBody := struct {
		Pid uint `json:"id"`
	}{}
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(200, models.Rep{
			State:   -1,
			Msg:     "请求参数有误",
			Data:    nil,
			Content: nil,
		})
		return
	}

	user, exists := c.Get("userinfo")
	if !exists || user.(models.User).WxOpenId != "fangxinadmin" {
		c.JSON(200, models.Rep{
			State:   -3,
			Msg:     "没有权限操作",
			Data:    nil,
			Content: nil,
		})
		return
	}

	tx := models.DB.Delete(&models.Product{}, reqBody.Pid)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State:   -2,
			Msg:     "删除失败, 请检查该产品是否存在",
			Data:    nil,
			Content: nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State:   0,
		Msg:     "删除成功",
		Data:    nil,
		Content: nil,
	})
	return

}

func (ProductController) AddProduct(c *gin.Context) {
	var Pro models.Product
	err := c.BindJSON(&Pro)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "添加失败, 请检查参数",
			Data:  nil,
		})
		return
	}
	tx := models.DB.Create(&Pro)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "添加失败",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "添加成功",
		Data:  Pro,
	})
	return

}

func (ProductController) ChangeProduct(c *gin.Context) {
	var Pro models.Product
	err := c.BindJSON(&Pro)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "更新失败, 请检查参数",
			Data:  nil,
		})
		return
	}
	tx := models.DB.Updates(Pro)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "更新失败",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "更新成功",
		Data:  nil,
	})
	return

}

func (ProductController) GetProductById(c *gin.Context) {
	var Pro models.Product
	err := c.BindJSON(&Pro)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "获取失败,参数错误",
			Data:  nil,
		})
		return
	}
	tx := models.DB.Find(&Pro)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "获取失败, 不存在该数据",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "获取成功",
		Data:  &Pro,
	})
}

// GetProductByName 通过产品名进行模糊搜索
func (ProductController) GetProductByName(c *gin.Context) {
	req := struct {
		Name string `json:"name"`
	}{}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "获取失败,请求参数错误",
			Data:  nil,
		})
		return
	}
	if req.Name == "" {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "拒绝空查询",
			Data:  nil,
		})
		return
	}
	var pros []models.Product
	tx := models.DB.Where("name LIKE ?", fmt.Sprint("%", req.Name, "%")).Find(&pros)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -3,
			Msg:   "查询失败",
			Data:  nil,
		})
		return
	}

	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "查询成功",
		Data:  pros,
	})
	return
}
