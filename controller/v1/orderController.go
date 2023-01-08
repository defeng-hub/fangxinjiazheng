package v1

import (
	"fangxinjiazheng/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct{}

func (OrderController) GetAllOrdersByUser(c *gin.Context) {
	userinfo, exits := c.Get("userinfo")
	if !exits {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "未登录, 请先登录",
			Data:  nil,
		})
		return
	}

	user := userinfo.(models.User)

	tx := models.DB.Preload("Orders", func(db *gorm.DB) *gorm.DB {
		return db.Preload("ProductObj")
	}).Find(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "获取失败",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "获取成功",
		Data:  user,
	})
	return
}

func (OrderController) AddOrder(c *gin.Context) {
	req := struct {
		Uid     uint   `json:"uid"`
		Pid     uint   `json:"pid"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		Payment int64  `json:"payment"`
	}{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "请求格式错误",
			Data:  nil,
		})
		return
	}
	if len(req.Name) == 0 || len(req.Address) == 0 || len(req.Phone) != 11 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "请求数据格式错误",
			Data:  nil,
		})
		return
	}
	user := models.User{ID: req.Uid}
	product := models.Product{ID: req.Pid}
	utx := models.DB.First(&user)
	if utx.Error != nil {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "不存在该用户,请检查用户id",
			Data:  nil,
		})
		return
	}
	ptx := models.DB.First(&product)
	if ptx.Error != nil {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "不存在该商品,请检查产品id",
			Data:  nil,
		})
		return
	}

	order := models.Order{
		Amount:    product.Price,
		Payment:   req.Payment,
		PayStatus: 0,
		Address:   req.Address,
		KdName:    req.Name,
		KdPhone:   req.Phone,
		ProductId: req.Pid,
		//ProductObj: product,
		UserId: req.Uid,
		//UserObj: user,
	}
	odx := models.DB.Create(&order)
	if odx.Error != nil || odx.RowsAffected == 0 {
		//fmt.Printf("%#v", odx.Error)
		c.JSON(200, models.Rep{
			State: -5,
			Msg:   "订单创建失败",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "订单创建成功",
		Data:  nil,
	})
	return

}

func (OrderController) GetAllOrder(c *gin.Context) {

	if userinfo, exits := c.Get("userinfo"); exits {
		user := userinfo.(models.User)
		if user.WxOpenId != "fangxinadmin" {
			c.JSON(200, models.Rep{
				State: -1,
				Msg:   "权限不足",
				Data:  nil,
			})
		}
	} else {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "未登录",
			Data:  nil,
		})
	}

	var orders []models.OrderAPI
	models.DB.Preload("ProductObj").Find(&orders)
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "获取成功",
		Data:  orders,
	})
}

func (OrderController) GetOrderById(c *gin.Context) {
	req := struct {
		ID uint `json:"id"`
	}{}

	err := c.BindJSON(&req)
	if err != nil || req.ID == 0 {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "请求数据格式错误",
			Data:  nil,
		})
		return
	}

	// 先获取登录订单, 在其中搜索 该订单id
	userinfo, exits := c.Get("userinfo")
	if !exits {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "未登录, 请先登录",
			Data:  nil,
		})
		return
	}

	user := userinfo.(models.User)
	order := models.Order{}
	tx := models.DB.Find(&order, "id = ? AND user_id = ?", req.ID, user.ID)

	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -3,
			Msg:   "查询用户订单失败,该订单id不属于当前登录用户,或者订单不存在",
			Data:  nil,
		})
		return
	}

	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "获取成功",
		Data:  order,
	})
	return
}
