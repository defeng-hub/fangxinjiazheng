package v1

import (
	"fangxinjiazheng/models"
	"fangxinjiazheng/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) LoginByAdmin(c *gin.Context) {
	var reqbody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//绑定json和结构体
	if err := c.BindJSON(&reqbody); err != nil {
		c.JSON(200, gin.H{"state": -1, "msg": "参数不正确"})
		return
	}
	//1 先执行登录逻辑
	if reqbody.Username == "admin" && reqbody.Password == "111" {
		//签发jwt
		token, err := util.GenerateToken("fangxinadmin")
		if err != nil {
			c.JSON(200, gin.H{"state": -2, "msg": "jwt签发失败"})
			return
		} else {
			c.JSON(200, gin.H{"state": 0, "token": token, "username": reqbody.Username, "msg": "签发成功"})
			return
		}

	} else {
		c.JSON(200, gin.H{"state": -1, "msg": "信息不正确"})
		return
	}

}

// LoginByWx 此处传入一个前端微信登录返回的code, 然后会执行注册或者登录
// 登录完成后会返回 jwt签名
func (UserController) LoginByWx(c *gin.Context) {
	str := "093TR20w3A7BJZ2Sdb2w3dpOAt0TR206"
	oauth2, err := util.WxSDK.GetWxOpenIdFromOauth2(str)
	if err != nil {
		c.JSON(200, models.Rep{
			State:   -1,
			Msg:     "登录失败,ErrCode:" + fmt.Sprintf("%+v", err),
			Data:    nil,
			Content: nil,
		})
		return
	}
	//fmt.Printf("%#v", oauth2)

	kuser := models.User{}
	tx := models.DB.First(&kuser, "wx_open_id = ?", oauth2.Openid)
	if tx.Error != nil || tx.RowsAffected == 0 {
		fmt.Println("没找到")
		user := models.User{
			Name:     "小程序用户",
			WxOpenId: oauth2.Openid,
			Phone:    "",
			Address:  "",
		}
		tx2 := models.DB.Create(&user)
		if tx2.Error == nil {
			token, err2 := util.GenerateToken(oauth2.Openid)
			if err2 == nil {
				c.JSON(200, gin.H{
					"state": 0,
					"token": token,
					"msg":   "首次登录,签发成功",
					"data":  user,
				})
				return
			}
		}
		c.JSON(200, models.Rep{
			State:   -5,
			Msg:     "登录失败",
			Data:    nil,
			Content: nil,
		})
		return
	} else {
		token, err2 := util.GenerateToken(oauth2.Openid)
		if err2 == nil {
			c.JSON(200, gin.H{
				"state": 0,
				"token": token,
				"msg":   "登录成功",
				"data":  kuser,
			})
			return
		}
		c.JSON(200, models.Rep{
			State:   -4,
			Msg:     "登录失败,ErrorCode:" + fmt.Sprintf("%+v", err2),
			Data:    nil,
			Content: nil,
		})
		return

	}
}

func (UserController) Logined(c *gin.Context) {
	user, exits := c.Get("userinfo")
	if !exits || user == nil {
		//fmt
		c.Abort()
		c.JSON(200, models.Rep{
			State:   -1,
			Msg:     "登录失败",
			Data:    nil,
			Content: nil,
		})
		return
	}
	m := map[string]string{}
	m["name"] = user.(models.User).Name
	c.JSON(200, models.Rep{
		State:   0,
		Msg:     "登录成功",
		Data:    m,
		Content: nil,
	})

}
