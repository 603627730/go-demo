package model

import "time"

type AdminUser struct {
	Id           int64     `gorm:"column:id;primary_key" json:"id"`           // ID
	CreateBy     string    `gorm:"column:create_by" json:"create_by"`         // 创建者
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`     // 创建时间
	UpdateBy     string    `gorm:"column:update_by" json:"update_by"`         // 更新者
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`     // 更新时间
	Avatar       string    `gorm:"column:avatar" json:"avatar"`               // 用户头像
	DepartmentId string    `gorm:"column:department_id" json:"department_id"` // 所属部门ID
	Description  string    `gorm:"column:description" json:"description"`     // 备注
	Email        string    `gorm:"column:email" json:"email"`                 // 邮件
	Mobile       string    `gorm:"column:mobile" json:"mobile"`               // 手机
	NickName     string    `gorm:"column:nick_name" json:"nick_name"`         // 昵称
	Password     string    `gorm:"column:password" json:"password"`           // 密码
	Sex          string    `gorm:"column:sex" json:"sex"`
	Username     string    `gorm:"column:username;NOT NULL" json:"username"` // 用户名
	RoleIds      string    `gorm:"column:role_ids" json:"role_ids"`          // 角色ID集合
}

func (m *AdminUser) TableName() string {
	return "`li_admin_user`"
}
