// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysAddonsType is the golang structure for table base_sys_addonsType.
type BaseSysAddonsType struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	Name       string      `json:"name"       orm:"name"       ` // 标题
	Image      string      `json:"image"      orm:"image"      ` // 图片
	Link       string      `json:"link"       orm:"link"       ` // 跳转
	Remark     string      `json:"remark"     orm:"remark"     ` // 备注
	Status     int         `json:"status"     orm:"status"     ` // 状态
	OrderNum   int         `json:"orderNum"   orm:"orderNum"   ` // 排序
}
