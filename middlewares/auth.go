package middlewares

import (
	"fangxinjiazheng/models"
	"fangxinjiazheng/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header["Authorization"]
		if len(authorization) == 0 {
			// 返回错误信息
			c.JSON(401, gin.H{"state": -1, "token": "token验证失败"})
			//终止程序
			c.Abort()
			return
		}
		if len(authorization[0]) <= 16 {
			// 返回错误信息
			c.JSON(401, gin.H{"state": -1, "token": "token验证失败"})
			//终止程序
			c.Abort()
			return
		}

		token := authorization[0][7:]
		claims, err := util.ParseToken(token)
		if err != nil {
			// 返回错误信息
			c.JSON(401, gin.H{"state": -1, "token": "token验证失败"})
			//终止程序
			c.Abort()
			return
		}

		//c.Set("userinfo", claims.Userinfo)
		//fmt.Println(claims.Userinfo)
		userinfo := models.User{}
		models.DB.Find(&userinfo, "wx_open_id = ?", claims.Userinfo)

		c.Set("userinfo", userinfo)

		c.Next()
	}
}
