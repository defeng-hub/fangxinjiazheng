package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint           `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `gorm:"comment:'产品名';" json:"name"`
	Price     float64        `gorm:"comment:'价格';" json:"price"`
	LunboImg1 string         `gorm:"comment:'轮播图1';" json:"lunboImg1"`
	LunboImg2 string         `gorm:"comment:'轮播图2';" json:"lunboImg2"`
	LunboImg3 string         `gorm:"comment:'轮播图3';" json:"lunboImg3"`
	LunboImg4 string         `gorm:"comment:'轮播图4';" json:"lunboImg4"`
	LongImg   string         `gorm:"comment:'长图';" json:"longImg"`

	Orders []Order `gorm:"foreignKey:ProductId;" json:"orders,omitempty"`
}

func (Product) TableName() string {
	return "fx_product"
}

type ProductAPI struct {
	ID        uint      `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `gorm:"comment:'产品名';" json:"name"`
	Price     float64   `gorm:"comment:'价格';" json:"price"`
}

func (ProductAPI) TableName() string {
	return "fx_product"
}
