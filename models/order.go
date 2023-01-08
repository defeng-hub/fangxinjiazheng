package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uint           `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Amount     float64 `gorm:"comment:总金额;default:0;" json:"amount"`
	Payment    int64   `gorm:"comment:支付方式,1.微信支付 2.支付宝支付;default:0;" json:"payment"`
	PayStatus  int64   `gorm:"comment:支付状态;" json:"paystatus"`
	Address    string  `gorm:"comment:收货地址;" json:"address"`
	KdName     string  `gorm:"comment:姓名;" json:"kdName"`
	KdPhone    string  `gorm:"comment:手机号;" json:"kdPhone"`
	ProductId  uint    `json:"ProductId"`
	ProductObj Product `gorm:"foreignKey:ProductId" json:"-"`
	UserId     uint    `json:"UserId"`
	UserObj    User    `gorm:"foreignKey:UserId" json:"-"`
}

func (Order) TableName() string {
	return "fx_order"
}

func (o Order) GetPayStatus() string {
	var status = map[int64]string{
		0: "未支付",
		1: "已付款",
		2: "已完结",
	}
	return status[o.PayStatus]
}

// OrderAPI OrderAPI是订单的展示类
type OrderAPI struct {
	ID         uint       `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	Amount     float64    `gorm:"comment:总金额;default:0;" json:"amount"`
	Payment    int64      `gorm:"comment:支付方式,1.微信支付 2.支付宝支付;default:0;" json:"payment"`
	PayStatus  int64      `gorm:"comment:支付状态;" json:"paystatus"`
	Address    string     `gorm:"comment:收货地址;" json:"address"`
	KdName     string     `gorm:"comment:姓名;" json:"kdName"`
	KdPhone    string     `gorm:"comment:手机号;" json:"kdPhone"`
	ProductId  uint       `json:"ProductId"`
	ProductObj ProductAPI `gorm:"foreignKey:ProductId" json:"product"`
	UserId     uint       `json:"UserId"`
}

func (OrderAPI) TableName() string {
	return "fx_order"
}
