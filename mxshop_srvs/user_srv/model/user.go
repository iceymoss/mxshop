package model

import (
	"time"

	"gorm.io/gorm"
)

//BaseModel 公共字段
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"` //column 定义别名
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

//User 定义用户信息
type User struct {
	BaseModel
	Mobile   string     `gorm:"idx_mobile;unique;type:varchar(11) comment 'idx_mobile表示索引用于快速查找';not null"`
	PassWord string     `gorm:"type:varchar(100) comment '加密后的密码';not null"`
	NickName string     `gorm:"type: varchar(20) comment '昵称' "`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'male表示男， famale表示女'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示用户， 2表示管理员'"`
}
