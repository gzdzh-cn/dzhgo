// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysRoleMenu is the golang structure of table base_sys_role_menu for DAO operations like Where/Data.
type BaseSysRoleMenu struct {
	g.Meta     `orm:"table:base_sys_role_menu, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	RoleId     interface{} //
	MenuId     interface{} //
}
