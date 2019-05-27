package model

import "time"

// Model common model
type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`        // id主键
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"` // 创建时间
	UpdatedAt time.Time  `json:"updated_at" gorm:"updated_at"` // 更新时间
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"` // 删除时间
}
