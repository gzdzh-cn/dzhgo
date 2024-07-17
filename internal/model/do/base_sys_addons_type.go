// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysAddonsType is the golang structure of table base_sys_addonsType for DAO operations like Where/Data.
type BaseSysAddonsType struct {
	g.Meta     `orm:"table:base_sys_addonsType, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	Name       interface{} // 标题
	Image      interface{} // 图片
	Link       interface{} // 跳转
	Remark     interface{} // 备注
	Status     interface{} // 状态
	OrderNum   interface{} // 排序
}
