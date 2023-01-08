package models

import (
	"gorm.io/gorm"
	"time"
)

type GlobalSetting struct {
}
type IndexAd struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Img       string         `json:"img"`
}

type Rep struct {
	State   int         `json:"state"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Content interface{} `json:"content,omitempty"`
}

type Img struct {
	ID        uint           `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Url       string         `json:"url"`
}

func (Img) TableName() string {
	return "fx_img"
}

func (IndexAd) TableName() string {
	return "fx_index_ad"
}
