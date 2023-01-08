package models

import (
	"gorm.io/gorm"
	"time"
)

// User 这是微信小程序的使用用户
type User struct {
	ID        uint           `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `gorm:"comment:'姓名';" json:"name,omitempty"`
	WxOpenId  string         `gorm:"comment:'微信唯一WxOpenId';" json:"-"`
	Phone     string         `gorm:"comment:'手机号';" json:"phone,omitempty"`
	Address   string         `gorm:"comment:'住址';" json:"address,omitempty"`
	Orders    []Order        `gorm:"foreignKey:UserId;" json:"orders"`
}

func (User) TableName() string {
	return "fx_user"
}

type Employee struct {
	ID        uint           `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `gorm:"comment:'姓名';" json:"name"`
	Age       string         `gorm:"comment:'年龄';" json:"age"`
	Sex       string         `gorm:"comment:'性别';" json:"sex"`
	Phone     string         `gorm:"comment:'手机号';" json:"phone"`
	Address   string         `gorm:"comment:'住址';" json:"address"`
}

func (Employee) TableName() string {
	return "fx_employee"
}
