// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysRoleMenuDao is the data access object for table base_sys_role_menu.
type BaseSysRoleMenuDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns BaseSysRoleMenuColumns // columns contains all the column names of Table for convenient usage.
}

// BaseSysRoleMenuColumns defines and stores column names for table base_sys_role_menu.
type BaseSysRoleMenuColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	RoleId     string //
	MenuId     string //
}

// baseSysRoleMenuColumns holds the columns for table base_sys_role_menu.
var baseSysRoleMenuColumns = BaseSysRoleMenuColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	RoleId:     "roleId",
	MenuId:     "menuId",
}

// NewBaseSysRoleMenuDao creates and returns a new DAO object for table data access.
func NewBaseSysRoleMenuDao() *BaseSysRoleMenuDao {
	return &BaseSysRoleMenuDao{
		group:   "default",
		table:   "base_sys_role_menu",
		columns: baseSysRoleMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseSysRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseSysRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseSysRoleMenuDao) Columns() BaseSysRoleMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseSysRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseSysRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseSysRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
