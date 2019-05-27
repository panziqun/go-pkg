package rbac

import (
	"github.com/laughmaker/go-pkg/model"
)

type Role struct {
	Name            string           `gorm:"column:name" json:"name"`
	RolePermissions []RolePermission `json:"role_permissions"`
	model.Model
}

type Permission struct {
	Name   string `gorm:"column:name" json:"name"`
	Method string `gorm:"column:method" json:"method"`
	Route  string `gorm:"column:route" json:"route"`
	model.Model
}

type RolePermission struct {
	RoleID       int        `gorm:"column:role_id" json:"role_id"`
	PermissionID int        `gorm:"column:permission_id" json:"permission_id"`
	Permission   Permission `json:"permission"`
	model.Model
}

type UserRole struct {
	UserID int  `gorm:"column:user_id" json:"user_id"`
	RoleID int  `gorm:"column:role_id" json:"role_id"`
	Role   Role `json:"role"`
	model.Model
}
