package model

import "time"

type Category struct {
	Id             int64     `gorm:"column:id;primary_key" json:"id"`               // ID
	CreateBy       string    `gorm:"column:create_by" json:"create_by"`             // 创建者
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`         // 创建时间
	UpdateBy       string    `gorm:"column:update_by" json:"update_by"`             // 更新者
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`         // 更新时间
	CommissionRate string    `gorm:"column:commission_rate" json:"commission_rate"` // 佣金比例
	Image          string    `gorm:"column:image" json:"image"`                     // 分类图标
	Level          int       `gorm:"column:level" json:"level"`                     // 层级
	Name           string    `gorm:"column:name" json:"name"`                       // 分类名称
	ParentId       int64     `gorm:"column:parent_id" json:"parent_id"`             // 父ID
	SortOrder      int       `gorm:"column:sort_order;default:0" json:"sort_order"` // 排序值
}

func (m *Category) TableName() string {
	return "`li_category`"
}
