package model

import (
	"time"

	"gorm.io/gorm"
)

//BaseModel 公共字段
type BaseModel struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"` //column 定义别名
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}
