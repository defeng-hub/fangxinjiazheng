package v1

import (
	"fangxinjiazheng/models"
	"fangxinjiazheng/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"path"
	"strconv"
	"time"
)

// EmployeeController 初始化一个 员工类用来保存 方法
type EmployeeController struct{}

// GetAllEmployee 获取全部员工
func (EmployeeController) GetAllEmployee(c *gin.Context) {
	var employees []models.Employee
	tx := models.DB.Find(&employees)
	if tx.Error != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "请求参数错误",
			Data:  nil,
		})
		return
	}

	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "获取成功",
		Data:  employees,
	})
	return
}

// AddEmployee 添加员工
func (EmployeeController) AddEmployee(c *gin.Context) {
	c.Get("userinfo")
	var employee models.Employee
	err := c.BindJSON(&employee)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "请求参数错误",
			Data:  nil,
		})
		return
	}
	if len(employee.Name) == 0 || len(employee.Phone) != 11 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "请检查参数格式",
			Data:  nil,
		})
		return
	}
	tx := models.DB.Create(&employee)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -3,
			Msg:   "新增员工失败失败",
			Data:  nil,
		})
		return
	}
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "添加成功",
		Data:  nil,
	})
	return

}

// DelEmployee 删除员工
func (EmployeeController) DelEmployee(c *gin.Context) {
	req := struct {
		Id uint `json:"id"`
	}{}
	err := c.ShouldBind(&req)
	if err != nil || req.Id == 0 {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "参数错误",
			Data:  nil,
		})
		return
	}

	tx := models.DB.Delete(&models.Employee{ID: req.Id})
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.JSON(200, models.Rep{
			State: -2,
			Msg:   "删除失败,请检查id是否存在",
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

// ChangeEmployee 修改员工信息接口
func (EmployeeController) ChangeEmployee(c *gin.Context) {
	var employee models.Employee
	err := c.BindJSON(&employee)
	if err != nil {
		c.JSON(200, models.Rep{
			State: -1,
			Msg:   "更新失败, 请检查参数",
			Data:  nil,
		})
		return
	}
	//此次updates只会更新非初始值
	tx := models.DB.Updates(employee)
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

// ExportExcel 导出excel功能
func (EmployeeController) ExportExcel(c *gin.Context) {
	var employees []models.Employee
	var obj models.Employee
	models.DB.Find(&employees)

	today := time.Now().Format("20060102")
	//创建多层夹
	err2 := os.MkdirAll("./static/excel/"+today, os.ModePerm)
	if err2 != nil {
		c.JSON(200, gin.H{
			"code": -3,
			"msg":  "保存失败",
		})
		return
	}
	filepath := path.Join("./static/excel/"+today, strconv.Itoa(int(time.Now().Unix()))+"-"+"export.xlsx")

	util.ExportDefaultExcel(obj, employees, filepath)
	config, _ := ini.Load("./conf/app.ini")
	api := config.Section("").Key("api").String()
	res := "http://" + api + "/" + filepath
	c.JSON(200, models.Rep{
		State: 0,
		Msg:   "导出成功",
		Data: map[string]interface{}{
			"url": res,
		},
	})
}
