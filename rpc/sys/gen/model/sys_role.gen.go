// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole 角色信息表
type SysRole struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                          // 编号
	RoleName   string     `gorm:"column:role_name;not null;comment:角色名称" json:"role_name"`                               // 角色名称
	RoleKey    string     `gorm:"column:role_key;not null;comment:权限字符" json:"role_key"`                                 // 权限字符
	RoleStatus int32      `gorm:"column:role_status;not null;default:1;comment:角色状态" json:"role_status"`                 // 角色状态
	RoleSort   int32      `gorm:"column:role_sort;not null;comment:角色排序" json:"role_sort"`                               // 角色排序
	DataScope  int32      `gorm:"column:data_scope;not null;comment:数据权限" json:"data_scope"`                             // 数据权限
	IsDeleted  int32      `gorm:"column:is_deleted;not null;comment:是否删除  0：否  1：是" json:"is_deleted"`                   // 是否删除  0：否  1：是
	IsAdmin    int32      `gorm:"column:is_admin;not null;comment:是否超级管理员:  0：否  1：是" json:"is_admin"`                   // 是否超级管理员:  0：否  1：是
	Remark     string     `gorm:"column:remark;not null;comment:备注" json:"remark"`                                       // 备注
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`                                // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
